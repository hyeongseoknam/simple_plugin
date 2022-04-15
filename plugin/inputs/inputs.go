package inputs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyeongseoknam/simple_plugin"

	"github.com/BurntSushi/toml"
)

type Creator func() simple_plugin.Input

var (
	Inputs = map[string]Creator{}
)

func Add(name string, creator Creator) {
	Inputs[name] = creator
}

func LoadPlugins(pluginPrefix string) []simple_plugin.Input {

	var allInputsThisTime []simple_plugin.Input
	listPluginConfigs(pluginPrefix, func(configfile string) {
		inputs, err := LoadConfig(&configfile)
		if err != nil {
			panic(err)
		}
		allInputsThisTime = append(allInputsThisTime, inputs...)
	})

	return allInputsThisTime
}

func LoadConfig(filePath *string) ([]simple_plugin.Input, error) {
	b, err := ioutil.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}

	conf := struct {
		Inputs map[string][]toml.Primitive
	}{}

	md, err := toml.Decode(string(b), &conf)
	if err != nil {
		return nil, err
	}

	fmt.Println("LoadConfig step -1", conf)
	return loadConfigIntoInputs(md, conf.Inputs)
}

func listPluginConfigs(pluginPrefix string, h1 func(string)) {
	err := filepath.Walk(pluginPrefix,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && info.Size() > 0 {
				fmt.Println(path, info.Name())

				h1(path)
			}

			return nil
		})
	if err != nil {
		panic(err)
	}
}

func loadConfigIntoInputs(md toml.MetaData, inputConfigs map[string][]toml.Primitive) ([]simple_plugin.Input, error) {
	renderedInputs := []simple_plugin.Input{}

	fmt.Println("loadConfigIntoInputs inputConfigs:", inputConfigs)
	for name, primitives := range inputConfigs {
		inputCreator, ok := Inputs[name]
		if !ok {
			return nil, errors.New("unknown input " + name)
		}

		for _, primitive := range primitives {
			inp := inputCreator()
			fmt.Println("loadConfigIntoInputs inp:", inp)
			if err := md.PrimitiveDecode(primitive, inp); err != nil {
				return nil, err
			}

			renderedInputs = append(renderedInputs, inp)
		}
	}
	return renderedInputs, nil
}

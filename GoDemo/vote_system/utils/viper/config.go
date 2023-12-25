// @Author zhangjiaozhu 2023/3/18 14:08:00
package viper

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err :=viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

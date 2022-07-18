import configparser
import os


class Config:

    def create_config(self, path):
        if not os.path.exists(path):
            config = configparser.ConfigParser()
            config.add_section("Settings")

            config.set("Settings", "autorun", "False")
            with open(path, "w") as file:
                config.write(file)
        else:
            return print("The file exists")

    def get_config(self, path):
        if not os.path.exists(path):
            self.create_config(path)

        config = configparser.ConfigParser()
        config.read(path)
        return config

    def get_setting(self, path, section, setting):
        config = self.get_config(path)
        value = config.get(section, setting)
        return value

    def update_setting(self, path, section, setting, value):
        config = self.get_config(path)
        config.set(section, setting, value)
        with open(path, "w") as file:
            config.write(file)

    def delete_setting(self, path, section, setting):
        config = self.get_config(path)
        config.remove_option(section, setting)
        with open(path, "w") as file:
            config.write(file)

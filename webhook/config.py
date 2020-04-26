import yaml

from model.config import Config, ServerConfig, PluginConfig

_config: Config = None


def parse_config(path: str):
    with open(path, "r") as f:
        config = yaml.safe_load(f)
    server_config = ServerConfig(host=config["server"]["host"], port=config["server"]["port"],
                                 debug=config["server"]["debug"])
    plugins_config = {}
    for k, v in config["plugins"].items():
        plugins_config[k] = PluginConfig(enabled=v["enabled"], args=v.get("args", {}))

    global _config
    _config = Config(version=config["version"], server_config=server_config,
                     plugin_config=plugins_config)
    return get_config()


def get_config() -> Config:
    return _config

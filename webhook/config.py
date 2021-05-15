import yaml

from model.config import Config, ServerConfig, PluginConfig

_config: Config = None


def parse_config(path: str):
    with open(path, "r") as f:
        config = yaml.safe_load(f)

    base_url = config["server"].get("base_url", "")
    if base_url == "":
        base_url = f"http://{config['server']['host']}:{config['server']['port']}"
    server_config = ServerConfig(host=config["server"]["host"], port=config["server"]["port"],
                                 debug=config["server"]["debug"],
                                 token=config["server"].get("token", ""),
                                 base_url=base_url)
    plugins_config = {}
    for k, v in config["plugins"].items():
        plugins_config[k] = PluginConfig(enabled=v["enabled"], args=v.get("args", {}))

    global _config
    _config = Config(version=config["version"], server_config=server_config,
                     plugin_config=plugins_config)
    return get_config()


def get_config() -> Config:
    return _config

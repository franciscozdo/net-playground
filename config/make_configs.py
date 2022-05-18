import yaml
import os

docker_conf = 'docker-compose.yaml'

def image_from_name(name):
    if name == 'master':
        return 'playground/master'
    return 'playground/host'

def make_docker(conf):
    machines = conf['machines']
    yml = ''
    yml += 'version: "3"\n'
    yml += '\n'

    yml += 'services:\n'
    for name in machines:
        yml += '  {}:\n'.format(name)
        yml += '    image: {}\n'.format(image_from_name(name))
        yml += '    container_name: {}\n'.format(name)
        if name == 'master':
            yml += '    stdin_open: true\n'
            yml += '    tty: true\n'
        yml += '    volumes:\n'
        yml += '      - ./config/{}:/var/app/data/config.yaml\n'.format(name + ".yaml")
        yml += '      - ./config/{}:/var/app/data/start.sh\n'.format(name + ".sh")
        yml += '    networks:\n'
        yml += '      fnet:\n'
        yml += '        ipv4_address: "{}"\n'.format(machines[name]['address'])
        yml += '\n'

    yml += 'networks:\n'
    yml += '  fnet:\n'
    yml += '    driver: bridge\n'
    yml += '    ipam:\n'
    yml += '      driver: default\n'
    yml += '      config:\n'
    yml += '        - subnet: "{}"\n'.format(conf['network'])

    return yml

def make_master(conf):
    yml = ''
    for m in conf:
        yml += '{}:\n'.format(m)
        for k in conf[m]:
            yml += '  {}: {}\n'.format(k, conf[m][k])

    f = open("master.yaml", "w")
    f.write(yml)
    f.close()

def make_host(name, conf):
    yml = ''
    yml += 'name: {}\n'.format(name)
    yml += 'port: {}\n'.format(conf[name]['port'])
    yml += 'connected:\n'
    for h in conf[name]['connected']:
        yml += '  - {}\n'.format(conf[h]['address'])
    yml += 'services:\n'
    for s in conf[name]['services']:
        yml += '  - {}\n'.format(s)

    fname = name + ".yaml"
    f = open(fname, "w")
    f.write(yml)
    f.close()

    exe = ''
    exe = 'ls /usr/local/bin\n'
    for s in conf[name]['services']:
        exe += 'nohup {} &\n'.format(s)
    exe += 'host -c /var/app/data/config.yaml\n'

    # create starting script
    fname = name + ".sh"
    f = open(fname, "w")
    f.write(exe)
    f.close()


def make_machines(conf):
    for name in conf:
        if name == 'master':
            make_master(conf)
        else:
            make_host(name, conf)


if __name__ == "__main__":
    with open("config.yaml") as f:
        data = yaml.load(f, Loader=yaml.loader.SafeLoader)

        df = open(docker_conf, "w")
        df.write(make_docker(data))
        df.close()
        make_machines(data['machines'])

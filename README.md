[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

# ansible-rackhd-inventory
A inventory generator script to work with RackHD [RackHD](https://rackhd.readthedocs.io/en/latest/). It will generate your groups by looking
at either from tagged nodes.

### Example

Set your configuration in `config.yaml` for `rackhd_api_url` and `groups`. After that just compile and run

```
   # compile
   $ go build -o inventory .

   # run
   $ ./inventory --list

```

and with ansible

```
    $ ansible -i inventory some-playbook.yml
```

### Output
```
{
    '_meta': {
        'hostvars': {
            'web-node1.dev': {
                'ansible_ssh_host': '127.0.0.1',
                'ansible_ssh_host_private': '127.0.0.1',
            },
            'web-node2.dev': {
                'ansible_ssh_host': '127.0.0.20',
                'ansible_ssh_host_private': '127.0.0.20',
            },
            'database1.dev': {
                'ansible_ssh_host': '192.168.0.1',
                'ansible_ssh_host_private': '192.168.0.1',
            },
            'database2.dev': {
                'ansible_ssh_host': '192.168.0.2',
                'ansible_ssh_host_private': '192.168.0.2',
            },
            'database3.dev': {
                'ansible_ssh_host': '192.168.0.4',
                'ansible_ssh_host_private': '192.168.0.4',
            },
        },
    },
    'web': {
        'hosts': [
            'web-node1.dev',
            'web-node2.dev',
        ]
    },
    'db': {
        'hosts': [
            'database1.dev',
            'database2.dev',
            'database3.dev',
        ]
    }
}
```

### Configuration

Available configuration:
- Config file `config.yaml`
```
rackhd_api_url: http://192.168.1.2:8080
# what groups do you want to create from your tags in rackhd ?
# a tag in rackhd will create a group
groups:
  - ceph-node
  - ceph-mon
```
- Environment vars:
    - `RACK_HD_API_URL` - Overrides the parameter set in `config.yaml`
    - `ANSIBLE_RACKHD_CONFIG_PATH` - Specifies the location of the config file
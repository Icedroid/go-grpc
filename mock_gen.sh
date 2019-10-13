#!/bin/bash

# dao
mockgen --source dao/playbook.go --destination dao/playbook_mock.go  --package dao --self_package git.fogcdn.top/axe/ops-playbook/dao
mockgen --source dao/playbook_file.go --destination dao/playbook_file_mock.go  --package dao --self_package git.fogcdn.top/axe/ops-playbook/dao
mockgen --source dao/playbook_entrypoint.go --destination dao/playbook_entrypoint_mock.go  --package dao --self_package git.fogcdn.top/axe/ops-playbook/dao
mockgen --source dao/template.go --destination dao/template_mock.go  --package dao --self_package git.fogcdn.top/axe/ops-playbook/dao

# move笔记

## 环境搭建

### 以下过程全程需要魔法（记得打开全局）

安装move开发环境需要安装python3.6+和rust以及cargo和git

笔者通过脚本安装aptos cli实现move环境的搭建

在powershell中输入以下命令

```
iwr "https://aptos.dev/scripts/install_cli.py" -useb | Select-Object -ExpandProperty Content | python3
```

如果收到错误，ModuleNotFoundError: No module named packaging

执行下面的命令

```
pip3 install packaging
```

笔者本人的python为3.11

执行下面的命令才安装成功

```
iwr "https://aptos.dev/scripts/install_cli.py" -useb | Select-Object -ExpandProperty Content | python
```

命令执行完之后

在cmd中输入aptos，如果出现以下的结果，则说明安装成功

![image-20240309233109129](assets/image-20240309233109129.png)

## 工具选择

vscode

插件

aptos-move-analyzer

CodeGeeX: AI Code AutoComplete, Chat, Auto Comment

## 初始化

1 win+r 输入cmd

2 cd进入队友文件夹 或者通过mkdir创建相应的文件夹之后cd进入

3 在终端中输入以下命令

```
aptos move init --name 你刚刚创建的文件夹的名字
```

如果成功，会出现以下结果

![img](assets/11VX]JD({W$%K565DQMH)E3.png)

4 完成后 在终端中输入  code .打开vscode

5 ctrl+esc下面的那个键唤出终端输入aptos init，然后一路enter向下就行 




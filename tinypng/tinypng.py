# -*- coding:utf-8 -*-
# !/usr/bin/env python3
# 使用tinypng API压缩项目图片
import tinify
import os
import time
import shutil

# tinify.key ="Your Developer API Key" # AppKey
# 设置代理
# tinify.proxy = ""

os.mkdir("/Output")

# 压缩图片的key
online_key_list = [
    "NtkTWRRNjqfF8qllPlrt6hMMHb3Wbb24", # 可以继续添加  防止一个key不够
]

# 获取key
online_key_list_iter = iter(online_key_list)
online_key = next(online_key_list_iter)

# 需要压缩图片的路径
fromPath = "/pic"  # source dir path
# 压缩后下载的图片路径
toPath = "/Output"  # dest dir path


tinifyAPi = tinify.tinify


# 在线压缩
def compress_online(sourcefile, outputfile):
    global online_key
    compresskey = online_key

    tinify.key = compresskey
    rs = False
    try:
        source = tinifyAPi.from_file(sourcefile)
        source.to_file(outputfile)
        print('压缩图片...' + outputfile)
        rs = True
        pass
    except tinify.AccountError:
        # Verify your API key and account limit.
        # 如果key值无效 换一个key继续压缩
        print("key值无效 换一个继续。。。")
        online_key = next(online_key_list_iter)
        compress_online(sourcefile, outputfile, name)  # 递归方法 继续读取
        rs = True

    except tinify.ClientError:
        # Check your source image and request options.
        print("Check your source image and request options.")
        rs = False
        pass
    except tinify.ServerError:
        # Temporary issue with the Tinify API.
        # print("Temporary issue with the Tinify API. %s" % e.message)
        print("Temporary issue with the Tinify API.")

        rs = False
        pass
    except tinify.ConnectionError:
        # A network connection error occurred.
        print("网络故障。。。休息1秒继续")
        time.sleep(1)
        compress_online(sourcefile, outputfile, name)  # 递归方法 继续读取
        rs = True
        pass
    except Exception:
        # Something else went wrong, unrelated to the Tinify API.
        print("Something else went wrong, unrelated to the Tinify API.")
        rs = False
        pass

    return rs


# root, dirs, files参数的含义：目录的路径(String)；root目录下所有子目录的名字(List)；root目录下非目录的名字
# os.walk(top,topdown=True,onerror=None)函数中各参数的含义：需要遍历的顶级目录的路径；默认值是“True”表示首先返回顶级目录下的文件，然后再遍历子目录中的文件；默认值为"None"，表示忽略文件遍历时的错误
for root, dirs, files in os.walk(fromPath):
    newToPath = toPath
    if len(root) > len(fromPath):  # 比较root和fromPath的字符长度
        innerPath = root[len(fromPath):]  # 字符串切割，将root字符串从第len(fromPath)个位置开始截取，不包括len(fromPath)这个位置的字符
        if innerPath[0] == '/':  # 判断innerPath的第一个字符是不是/符号
            innerPath = innerPath[1:]  # 字符串切割，例如innerPath的值为\test，那么innerPath[1:]之后的值为test
        newToPath = os.path.join(toPath, innerPath)  # 将toPath目录和innerPath文件或文件夹拼接之后的路径赋值给newToPath

    for name in files:
        newFromFilePath = os.path.join(root, name)
        newToFilePath = os.path.join(newToPath, name)
        fileName, fileSuffix = os.path.splitext(name)  # 分解文件名的扩展名
        if fileSuffix == '.png' or fileSuffix == '.jpg' or fileSuffix == '.jpeg' or fileSuffix == '.PNG' or fileSuffix == '.JPG' or fileSuffix == '.JPEG':
            # 			source = tinify.from_file(newFromFilePath)
            # 			source.to_file(newToFilePath)
            # 在线压缩
            if not compress_online(newFromFilePath, newToFilePath):
                print("压缩失败，检查报错信息")
                exit()
                pass
        else:
            pass

    for dirName in dirs:
        if os.path.exists(os.path.join(newToPath, dirName)):
            pass
        else:
            os.mkdir(os.path.join(newToPath, dirName))

shutil.move("/Output","/pic")

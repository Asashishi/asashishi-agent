#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
测试交互式程序
"""
import subprocess
import sys
import time

def test_interactive():
    """测试交互式程序"""
    print("🚀 开始测试交互式程序...")
    
    # 启动主程序
    proc = subprocess.Popen(
        ['python3', 'main.py'],
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        text=True,
        bufsize=1
    )
    
    # 等待程序启动
    time.sleep(0.5)
    
    # 发送测试输入序列
    inputs = [
        '1\n',      # 选择加法
        '10\n',     # 第一个数字
        '5\n',      # 第二个数字
        '0\n',      # 退出程序
    ]
    
    print("📤 发送测试输入...")
    for input_str in inputs:
        proc.stdin.write(input_str)
        proc.stdin.flush()
        time.sleep(0.2)
    
    # 获取输出
    stdout, stderr = proc.communicate(timeout=5)
    
    print("📥 程序输出:")
    print("="*50)
    print(stdout)
    print("="*50)
    
    if stderr:
        print("❌ 错误输出:")
        print(stderr)
    
    print(f"📊 退出码: {proc.returncode}")
    
    # 检查日志文件
    import os
    log_dir = "log"
    if os.path.exists(log_dir):
        log_files = [f for f in os.listdir(log_dir) if f.endswith('.md')]
        if log_files:
            latest_log = os.path.join(log_dir, sorted(log_files)[-1])
            print(f"📝 最新日志文件: {latest_log}")
            with open(latest_log, 'r', encoding='utf-8') as f:
                content = f.read()
                if "加法: 10.0 + 5.0" in content:
                    print("✅ 交互测试成功！操作已记录到日志")
                else:
                    print("⚠️ 日志中未找到测试操作记录")

if __name__ == "__main__":
    test_interactive()

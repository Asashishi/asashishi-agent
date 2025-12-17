
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Python 演示项目主程序
这是一个简单的交互式计算器程序
"""

import sys
from typing import NoReturn
from datetime import datetime

# 设置 UTF-8 编码
sys.stdout.reconfigure(encoding='utf-8')


def setup_logging() -> None:
    """设置日志系统"""
    import os
    log_dir: str = "log"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)
    
    timestamp: str = datetime.now().strftime("%Y%m%d%H%M%S")
    log_file: str = os.path.join(log_dir, f"{timestamp}.md")
    
    with open(log_file, 'w', encoding='utf-8') as f:
        f.write(f"# 程序运行日志 - {timestamp}\n\n")
        f.write("## 程序启动\n")
        f.write(f"- 启动时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")
        f.write(f"- Python 版本: {sys.version}\n\n")
    
    print(f"📝 日志已初始化: {log_file}")


def log_operation(operation: str, result: str) -> None:
    """记录操作日志"""
    import os
    log_dir: str = "log"
    log_files: list[str] = [f for f in os.listdir(log_dir) if f.endswith('.md')]
    
    if log_files:
        latest_log: str = os.path.join(log_dir, sorted(log_files)[-1])
        with open(latest_log, 'a', encoding='utf-8') as f:
            f.write(f"## {datetime.now().strftime('%H:%M:%S')}\n")
            f.write(f"- 操作: {operation}\n")
            f.write(f"- 结果: {result}\n\n")


def display_menu() -> None:
    """显示主菜单"""
    print("\n" + "="*50)
    print("🐍 Python 演示计算器 🐍")
    print("="*50)
    print("1. 加法运算")
    print("2. 减法运算")
    print("3. 乘法运算")
    print("4. 除法运算")
    print("5. 显示帮助")
    print("0. 退出程序")
    print("="*50)


def get_number_input(prompt: str) -> float:
    """获取数字输入"""
    while True:
        try:
            value: str = input(prompt)
            return float(value)
        except ValueError:
            print("❌ 请输入有效的数字！")


def main() -> NoReturn:
    """主函数"""
    print("🚀 Python 演示项目启动中...")
    setup_logging()
    
    # 动态导入模块
    try:
        from src.calculator import Calculator
        from src.utils import format_result, validate_division
    except ImportError as e:
        print(f"❌ 导入模块失败: {e}")
        print("请确保项目结构正确，src 目录包含必要的模块")
        sys.exit(1)
    
    # 创建计算器实例
    calculator: Calculator = Calculator()
    
    while True:
        display_menu()
        choice: str = input("请选择操作 (0-5): ").strip()
        
        if choice == '0':
            print("\n👋 感谢使用，再见！")
            log_operation("程序退出", "用户选择退出")
            sys.exit(0)
        
        elif choice == '1':
            print("\n➕ 加法运算")
            num1: float = get_number_input("请输入第一个数字: ")
            num2: float = get_number_input("请输入第二个数字: ")
            result: float = calculator.add(num1, num2)
            formatted: str = format_result("加法", num1, num2, result)
            print(f"✅ 结果: {formatted}")
            log_operation(f"加法: {num1} + {num2}", f"结果: {result}")
        
        elif choice == '2':
            print("\n➖ 减法运算")
            num1: float = get_number_input("请输入第一个数字: ")
            num2: float = get_number_input("请输入第二个数字: ")
            result: float = calculator.subtract(num1, num2)
            formatted: str = format_result("减法", num1, num2, result)
            print(f"✅ 结果: {formatted}")
            log_operation(f"减法: {num1} - {num2}", f"结果: {result}")
        
        elif choice == '3':
            print("\n✖️ 乘法运算")
            num1: float = get_number_input("请输入第一个数字: ")
            num2: float = get_number_input("请输入第二个数字: ")
            result: float = calculator.multiply(num1, num2)
            formatted: str = format_result("乘法", num1, num2, result)
            print(f"✅ 结果: {formatted}")
            log_operation(f"乘法: {num1} × {num2}", f"结果: {result}")
        
        elif choice == '4':
            print("\n➗ 除法运算")
            num1: float = get_number_input("请输入第一个数字: ")
            num2: float = get_number_input("请输入第二个数字: ")
            
            if validate_division(num2):
                result: float = calculator.divide(num1, num2)
                formatted: str = format_result("除法", num1, num2, result)
                print(f"✅ 结果: {formatted}")
                log_operation(f"除法: {num1} ÷ {num2}", f"结果: {result}")
            else:
                print("❌ 错误：除数不能为零！")
                log_operation(f"除法: {num1} ÷ {num2}", "错误：除数不能为零")
        
        elif choice == '5':
            print("\n📚 帮助信息")
            print("这是一个简单的交互式计算器程序")
            print("支持基本的四则运算")
            print("所有操作都会被记录到日志文件中")
            print("输入 0 可以退出程序")
            log_operation("查看帮助", "显示帮助信息")
        
        else:
            print("❌ 无效的选择，请重新输入！")
            log_operation(f"无效选择: {choice}", "用户输入无效选项")


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\n\n⚠️ 程序被用户中断")
        log_operation("程序中断", "用户按下 Ctrl+C")
        sys.exit(0)
    except Exception as e:
        print(f"\n❌ 程序发生错误: {e}")
        log_operation("程序错误", f"异常: {str(e)}")
        sys.exit(1)
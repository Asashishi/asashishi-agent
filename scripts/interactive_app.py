
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
交互式脚本程序示例
一个简单的交互式应用程序，演示用户输入处理
"""

import sys
from typing import List, Dict, Any, Optional


def setup_encoding() -> None:
    """设置系统编码为 UTF-8"""
    sys.stdout.reconfigure(encoding='utf-8')
    print("✅ 系统编码已设置为 UTF-8")


def display_menu() -> None:
    """显示主菜单"""
    print("\n" + "=" * 50)
    print("🎮 交互式脚本程序")
    print("=" * 50)
    print("1. 📝 输入并显示文本")
    print("2. 🔢 数字计算器")
    print("3. 📊 数据统计")
    print("4. 🎲 随机游戏")
    print("5. ℹ️  程序信息")
    print("0. 🚪 退出程序")
    print("=" * 50)


def handle_text_input() -> None:
    """处理文本输入功能"""
    print("\n📝 文本输入功能")
    print("请输入一些文本（输入 'done' 结束）：")
    
    lines: List[str] = []
    while True:
        user_input: str = input("> ").strip()
        if user_input.lower() == 'done':
            break
        if user_input:
            lines.append(user_input)
    
    if lines:
        print(f"\n📄 您输入了 {len(lines)} 行文本：")
        for i, line in enumerate(lines, 1):
            print(f"  {i}. {line}")
    else:
        print("⚠️  没有输入任何文本")


def handle_calculator() -> None:
    """处理计算器功能"""
    print("\n🔢 简单计算器")
    
    try:
        num1: float = float(input("请输入第一个数字: "))
        num2: float = float(input("请输入第二个数字: "))
        operation: str = input("请输入操作符 (+, -, *, /): ").strip()
        
        result: Optional[float] = None
        if operation == '+':
            result = num1 + num2
        elif operation == '-':
            result = num1 - num2
        elif operation == '*':
            result = num1 * num2
        elif operation == '/':
            if num2 == 0:
                print("❌ 错误：除数不能为零")
                return
            result = num1 / num2
        else:
            print(f"❌ 不支持的操作符: {operation}")
            return
        
        print(f"✅ 计算结果: {num1} {operation} {num2} = {result}")
        
    except ValueError:
        print("❌ 错误：请输入有效的数字")


def handle_statistics() -> None:
    """处理数据统计功能"""
    print("\n📊 数据统计")
    print("请输入一系列数字，用空格分隔：")
    
    try:
        input_str: str = input("> ").strip()
        if not input_str:
            print("⚠️  没有输入任何数字")
            return
        
        numbers: List[float] = [float(x) for x in input_str.split()]
        
        if not numbers:
            print("⚠️  没有有效的数字")
            return
        
        count: int = len(numbers)
        total: float = sum(numbers)
        average: float = total / count
        maximum: float = max(numbers)
        minimum: float = min(numbers)
        
        print(f"\n📈 统计结果：")
        print(f"  数量: {count}")
        print(f"  总和: {total}")
        print(f"  平均值: {average:.2f}")
        print(f"  最大值: {maximum}")
        print(f"  最小值: {minimum}")
        
    except ValueError:
        print("❌ 错误：请输入有效的数字")


def handle_game() -> None:
    """处理随机游戏功能"""
    import random
    
    print("\n🎲 猜数字游戏")
    print("我已经想好了一个 1-100 之间的数字，试试猜猜看！")
    
    secret_number: int = random.randint(1, 100)
    attempts: int = 0
    max_attempts: int = 10
    
    while attempts < max_attempts:
        try:
            guess: int = int(input(f"第 {attempts + 1} 次尝试 (1-100): "))
            
            if guess < 1 or guess > 100:
                print("⚠️  请输入 1-100 之间的数字")
                continue
            
            attempts += 1
            
            if guess < secret_number:
                print("📈 猜小了！")
            elif guess > secret_number:
                print("📉 猜大了！")
            else:
                print(f"🎉 恭喜！你在第 {attempts} 次猜中了！")
                return
            
            if attempts < max_attempts:
                print(f"还有 {max_attempts - attempts} 次机会")
                
        except ValueError:
            print("❌ 错误：请输入有效的数字")
    
    print(f"😔 游戏结束！正确答案是: {secret_number}")


def show_program_info() -> None:
    """显示程序信息"""
    print("\nℹ️  程序信息")
    print("程序名称: 交互式脚本程序")
    print("版本: 1.0.0")
    print("作者: AI 代码助手")
    print("描述: 一个演示交互式编程的示例程序")
    print("特点:")
    print("  - 支持多种交互功能")
    print("  - 提供明确的退出选项")
    print("  - 包含错误处理机制")
    print("  - 符合 Python 编码规范")


def main() -> None:
    """主函数"""
    setup_encoding()
    
    print("✨ 欢迎使用交互式脚本程序！")
    print("这是一个演示用户交互的示例程序。")
    
    while True:
        display_menu()
        
        try:
            choice: str = input("请选择功能 (0-5): ").strip()
            
            if choice == '0':
                print("\n👋 感谢使用，再见！")
                break
            elif choice == '1':
                handle_text_input()
            elif choice == '2':
                handle_calculator()
            elif choice == '3':
                handle_statistics()
            elif choice == '4':
                handle_game()
            elif choice == '5':
                show_program_info()
            else:
                print(f"❌ 无效选择: {choice}，请重新输入")
                
        except KeyboardInterrupt:
            print("\n⚠️  检测到中断信号，请使用菜单选项 0 退出程序")
        except EOFError:
            print("\n⚠️  检测到文件结束，请使用菜单选项 0 退出程序")
        
        input("\n按 Enter 键继续...")


if __name__ == "__main__":
    try:
        main()
    except Exception as e:
        print(f"❌ 程序发生错误: {e}")
        sys.exit(1)
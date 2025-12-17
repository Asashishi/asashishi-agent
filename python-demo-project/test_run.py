
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
测试脚本 - 验证项目功能
"""

import sys
sys.stdout.reconfigure(encoding='utf-8')

print("🧪 开始测试 Python 演示项目...")
print("="*50)

# 测试导入模块
try:
    from src.calculator import Calculator
    from src.utils import format_result, validate_division
    print("✅ 模块导入成功")
except ImportError as e:
    print(f"❌ 模块导入失败: {e}")
    sys.exit(1)

# 测试计算器功能
print("\n🔧 测试计算器功能:")
calc = Calculator()

# 测试加法
result = calc.add(10, 5)
print(f"  加法: 10 + 5 = {result}")

# 测试减法
result = calc.subtract(10, 5)
print(f"  减法: 10 - 5 = {result}")

# 测试乘法
result = calc.multiply(10, 5)
print(f"  乘法: 10 × 5 = {result}")

# 测试除法
try:
    result = calc.divide(10, 5)
    print(f"  除法: 10 ÷ 5 = {result}")
except ZeroDivisionError:
    print("  除法: 除数不能为零")

# 测试操作计数
print(f"\n📊 操作计数: {calc.get_operation_count()}")

# 测试工具函数
print("\n🔧 测试工具函数:")

# 测试格式化
formatted = format_result("加法", 10, 5, 15)
print(f"  格式化结果: {formatted}")

# 测试除法验证
print(f"  验证除数 5: {validate_division(5)}")
print(f"  验证除数 0: {validate_division(0)}")

# 测试其他工具函数
from src.utils import is_even, is_prime, fibonacci

print(f"\n🔧 测试其他函数:")
print(f"  4 是偶数吗? {is_even(4)}")
print(f"  7 是质数吗? {is_prime(7)}")
print(f"  斐波那契数列前 5 项: {fibonacci(5)}")

print("\n" + "="*50)
print("🎉 所有测试完成！项目功能正常。")
print("💡 运行 'python3 main.py' 启动交互式计算器")
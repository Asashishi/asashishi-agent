
# Python 演示项目 🐍

这是一个简单的 Python 演示项目，展示了基本的项目结构和代码组织。

## 项目结构 📁

```
python-demo-project/
├── src/                    # 源代码目录
│   ├── __init__.py
│   ├── calculator.py      # 计算器模块
│   └── utils.py          # 工具函数模块
├── tests/                 # 测试目录
│   ├── __init__.py
│   ├── test_calculator.py
│   └── test_utils.py
├── docs/                  # 文档目录
├── log/                   # 日志目录
├── main.py               # 主程序入口
├── requirements.txt      # 项目依赖
└── README.md            # 项目说明
```

## 功能特性 ✨

1. **计算器模块**：提供基本的数学运算功能
2. **工具函数**：包含常用的工具函数
3. **日志记录**：自动记录程序运行日志
4. **单元测试**：包含完整的测试用例

## 快速开始 🚀

### 安装依赖
```bash
# 此项目不需要额外依赖
```

### 运行程序
```bash
python main.py
```

### 运行测试
```bash
python -m pytest tests/
```

## 使用示例 📝

```python
from src.calculator import Calculator
from src.utils import format_result

# 创建计算器实例
calc = Calculator()

# 执行计算
result = calc.add(10, 5)
formatted = format_result("加法", 10, 5, result)
print(formatted)
```

## 许可证 📄

MIT License
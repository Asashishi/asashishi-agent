
"""
Python 演示项目 - 源代码包
"""

__version__ = "1.0.0"
__author__ = "Python Demo Project"
__description__ = "一个简单的 Python 演示项目"

# 导出公共接口
from .calculator import Calculator
from .utils import format_result, validate_division

__all__ = [
    'Calculator',
    'format_result',
    'validate_division'
]
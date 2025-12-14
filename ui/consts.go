package ui

// Reset all styles
const Reset string = "\x1b[0m"

const FgColorBase string = "\x1b[38;" // 前景色
const BgColorBase string = "\x1b[48;" // 背景色

// Text styles
const Bold string = "\x1b[1m"          // 粗体
const Dim string = "\x1b[2m"           // 弱化
const Italic string = "\x1b[3m"        // 斜体
const Underline string = "\x1b[4m"     // 下划线
const Blink string = "\x1b[5m"         // 慢闪烁
const FastBlink string = "\x1b[6m"     // 快闪烁（大部分终端不支持）
const Reverse string = "\x1b[7m"       // 反色（前景/背景互换）
const Hidden string = "\x1b[8m"        // 隐藏
const Strikethrough string = "\x1b[9m" // 删除线

const ExceptionAtColorHex string = "Incorrect color hex"

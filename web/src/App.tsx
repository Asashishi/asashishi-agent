import { useState, type ChangeEvent, type JSX } from 'react'
import type { ContextStorageItem, WebSocketMsg } from './types';
import { MainPageContext } from './context';

// 无论使用框架还是原生 实现 websocket 回调和双向绑定即可完成实时数据交互和展示
let socket: WebSocket = new WebSocket("ws://localhost:3000/ws");

const displayMsg = (data: WebSocketMsg): void => {
  if (data.type === "exec_output") {
    const scpOutputState: ContextStorageItem<string> | undefined = MainPageContext.get<string>("scpOutput");
    scpOutputState?.setValue(scpOutputState.value + data.content + "\n");
  } else {
    const aiOutputState: ContextStorageItem<string> | undefined = MainPageContext.get<string>("aiOutput");
    aiOutputState?.setValue(aiOutputState.value + data.content);
  }
}

socket.onopen = () => {
  console.log("✅ 已连接到 WebSocket 服务");
};
socket.onmessage = function(event: MessageEvent<string>) {
  const tAraeState: ContextStorageItem<string> | undefined = MainPageContext.get<string>("tAraeValue");
  tAraeState?.setValue("");
  displayMsg(JSON.parse(event.data));
};
socket.onclose = function() {
  console.log("\n ❌ 连接已关闭");
};
socket.onerror = function(error) {
  console.error("\n ⚠️ 出错: " + error);
  // 如果是无关紧要的出错(如网络) 重连可以解决问题 服务器仅维护一个 socket 对象 不支持多个连接 如果需要可以自己改写
  socket = new WebSocket("ws://localhost:3000/ws");
};

const App: React.FC = (): JSX.Element => {
  const [aiOutput, setAiOutput] = useState<string>("");
  const [scpOutput, setScpOutput] = useState<string>("");
  const [tAraeValue, setTAraeValue] = useState<string>("");

  MainPageContext.set<string>([
    {
      key: "aiOutput",
      contextItem: {
        value: aiOutput,
        setValue: setAiOutput,
      },
    },
    {
      key: "scpOutput",
      contextItem: {
        value: scpOutput,
        setValue: setScpOutput,
      }
    },
    {
      key: "tAraeValue",
      contextItem: {
        value: tAraeValue,
        setValue: setTAraeValue,
      }
    }
  ]);
  return (
    <div>
      <h2>WebSocket Example</h2>
      <div>
        <label htmlFor="uInput">Input: </label>
        <div>
          <textarea
            id="uInput"
            rows={10}
            cols={50}
            value={tAraeValue}
            onChange={(event: ChangeEvent<HTMLTextAreaElement>) => setTAraeValue(event.target.value)}
          />
        </div>
        <br />
        <button onClick={() => socket.send(
          JSON.stringify({
            type: 'user_input',
            content: tAraeValue,
          }),
        )}>Send</button>
      </div>
      <br />
      <label>AI: </label>
      <button onClick={() => setAiOutput("")}>Clear</button>
      <div id="ai">{aiOutput}</div>
      <br />
      <label>SCP: </label>
      <button onClick={() => setScpOutput("")}>Clear</button>
      <div id="scp">{scpOutput}</div>
    </div>
  );
}

export default App;

import { useState, type ChangeEvent, type JSX } from 'react'
import { MainPageContext } from './context';
import AsashishiAgentWs from './utils/websocket';

const wsInstance: AsashishiAgentWs = new AsashishiAgentWs();

const injectContextItems = (): void => {
  wsInstance.injectContextItems(
    MainPageContext.get("tAraeValue")!,
    MainPageContext.get("aiOutput")!,
    MainPageContext.get("scpOutput")!,
  );
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

  injectContextItems();

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
        <button onClick={() => wsInstance.send(
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

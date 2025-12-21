import { useState, type ChangeEvent, type JSX } from 'react'

const App: React.FC = (): JSX.Element => {
  const [tAraeValue, setTAraeValue] = useState<string>("");
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
        <button onClick={() => {}}>Send</button>
      </div>
      <br />
      <label>AI: </label>
      <button onClick={() => {}}>Clear</button>
      <div id="ai">{}</div>
      <br />
      <label>SCP: </label>
      <button onClick={() => {}}>Clear</button>
      <div id="scp">{}</div>
    </div>
  );
}

export default App;

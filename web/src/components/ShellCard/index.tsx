import React, {
    useState,
    useContext
} from 'react';
import type { JSX } from 'react';
import wsInstance from '../../utils/websocket';
import { AsashishiAgentContext } from '../../context';

const ShellCard: React.FC = (): JSX.Element => {
    const { tab } = useContext(AsashishiAgentContext);
    const [shellOutput, setShellOutput] = useState<string>("");
    wsInstance.injectDependencies([{
        key: "shellOutput",
        value: shellOutput,
        setValue: setShellOutput,
    }]);
    return (
        <div style={{ display: tab === "Shell" ? "" : "none"  }}>
            <label>Shell: </label>
            <button onClick={() => setShellOutput("")}>Clear</button>
            <div style={{ whiteSpace: "pre-wrap" }}>{shellOutput}</div>
        </div>
    );
};

export default ShellCard;
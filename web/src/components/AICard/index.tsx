import {
    useContext,
    useState,
} from 'react';
import type { JSX } from 'react';
import wsInstance from '../../utils/websocket';
import { AsashishiAgentContext } from '../../context';

const AICard: React.FC = (): JSX.Element => {
    const { tab } = useContext(AsashishiAgentContext);
    const [aiOutput, setAIOutput] = useState("");
    wsInstance.injectDependencies([{
        key: "aiOutput",
        value: aiOutput,
        setValue: setAIOutput,
    }]);

    return (
        <div style={{ display: tab === "AI" ? "" : "none" }}>
            <label>AI: </label>
            <button onClick={() => setAIOutput("")}>Clear</button>
            <div id="ai">{aiOutput}</div>
        </div>
    );
};

export default AICard;
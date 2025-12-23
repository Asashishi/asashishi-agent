import {
    useContext,
    useState,
} from 'react';
import type { JSX } from 'react';
import wsInstance from '../../utils/websocket';
import { AsashishiAgentContext } from '../../context';
import styles from "./index.module.css";

const AICard: React.FC = (): JSX.Element => {

    const { tab } = useContext(AsashishiAgentContext);
    const [aiOutput, setAIOutput] = useState<string>("");
    const [aiOutputHistories, setAIOutputHistories] = useState<string[]>([]);
    wsInstance.injectDependencies([
        {
            key: "aiOutput",
            value: aiOutput,
            setValue: setAIOutput,
        },
        {
            key: "aiOutputHistories",
            value: aiOutputHistories,
            setValue: setAIOutputHistories,
        },
    ]);
    return (
        <div style={{ display: tab === "AI" ? "" : "none" }} className={styles.AIOutputWrapper}>
            {aiOutputHistories.map((item: string): JSX.Element => <div className={styles.AIOutput}>{item}</div>)}
            <div className={styles.AIOutput}>{aiOutput}</div>
        </div>
    );
};

export default AICard;
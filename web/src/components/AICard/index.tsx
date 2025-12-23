import {
    useContext,
    useState,
} from 'react';
import type { JSX } from 'react';
import wsInstance from '../../utils/websocket';
import { AsashishiAgentContext } from '../../context';
import styles from "./index.module.css";
import type { DisplayMsg } from '../../types/websocket_type';

const AICard: React.FC = (): JSX.Element => {

    const { tab } = useContext(AsashishiAgentContext);
    const [aiOutput, setAIOutput] = useState<string>("");
    const [outputHistories, setoutputHistories] = useState<DisplayMsg[]>([]);
    wsInstance.injectDependencies([
        {
            key: "aiOutput",
            value: aiOutput,
            setValue: setAIOutput,
        },
        {
            key: "outputHistories",
            value: outputHistories,
            setValue: setoutputHistories,
        },
    ]);
    return (
        <div style={{ display: tab === "AI" ? "" : "none" }} className={styles.AIOutputWrapper}>
            {outputHistories.map((item: DisplayMsg): JSX.Element => {
                if (item.type === "input") {
                    return (
                        <div className={styles.UserInput}>
                            <span>
                                User
                            </span>
                            {item.content}
                        </div>
                    )
                }
                return <div className={styles.AIOutput}>{item.content}</div>
            })}
            {aiOutput && <div className={styles.AIOutput}>{aiOutput}</div>}
        </div>
    );
};

export default AICard;
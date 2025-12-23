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
    const { ioHistories, setIOHistories } = useContext(AsashishiAgentContext);
    wsInstance.injectDependencies([
        {
            key: "aiOutput",
            value: aiOutput,
            setValue: setAIOutput,
        },
        {
            key: "ioHistories",
            value: ioHistories,
            setValue: setIOHistories,
        },
    ]);
    return (
        <div style={{ display: tab === "AI" ? "" : "none" }} className={styles.AIOutputWrapper}>
            {ioHistories.map((item: DisplayMsg): JSX.Element => {
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
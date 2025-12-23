import React, {
    useContext,
    useEffect,
    useRef,
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
    const cardRef: React.Ref<HTMLDivElement> | undefined = useRef(null);
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
    useEffect(() => {
        if (cardRef.current) {
            cardRef.current?.scrollTo({
              top: cardRef.current.scrollHeight,
              behavior: "smooth",
            });
        }
    }, [aiOutput]);
    return (
        <div
            ref={cardRef}
            className={styles.AIOutputWrapper}
            style={{ display: tab === "chat" ? "" : "none" }}
        >
            {ioHistories.map((item: DisplayMsg): JSX.Element | void => {
                if (item.type === "input" && item.diplayPosition === "chat") {
                    return (
                        <div className={styles.UserInput}>
                            <span>
                                User
                            </span>
                            {item.content}
                        </div>
                    )
                } else if (item.type === "output") {
                    return <div className={styles.AIOutput}>{item.content}</div>
                }
            })}
            {aiOutput && <div className={styles.AIOutput}>{aiOutput}</div>}
        </div>
    );
};

export default AICard;
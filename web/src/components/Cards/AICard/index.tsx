import React, {
    useContext,
    useEffect,
    useRef,
    useState,
} from 'react';
import type { JSX } from 'react';
import wsInstance from '../../../utils/websocket';
import { AsashishiAgentContext } from '../../../context';
import type { DisplayMsg } from '../../../types/websocket_type';
import styles from "../index.module.css";

const AICard: React.FC = (): JSX.Element => {
    const [aiOutput, setAIOutput] = useState<string>("");
    const { tab, ioHistories } = useContext(AsashishiAgentContext);
    const cardRef: React.Ref<HTMLDivElement> | undefined = useRef(null);
    wsInstance.injectDependencies([
        {
            key: "aiOutput",
            value: aiOutput,
            setValue: setAIOutput,
        }
    ]);
    useEffect(() => {
        if (cardRef.current) {
            cardRef.current?.scrollTo({
              top: cardRef.current.scrollHeight,
              behavior: "smooth",
            });
        }
        if (!ioHistories.length) {
            setAIOutput("");
        }
    }, [ioHistories, aiOutput]);
    return (
        <div
            ref={cardRef}
            className={styles.IOOutputWrapper}
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
                    return <div className={styles.IOOutput}>{item.content}</div>
                }
            })}
            {aiOutput && <div className={styles.IOOutput}>{aiOutput}</div>}
        </div>
    );
};

export default AICard;
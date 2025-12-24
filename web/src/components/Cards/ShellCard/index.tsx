import React, {
    useState,
    useContext,
    useRef,
    useEffect
} from 'react';
import type { JSX } from 'react';
import wsInstance from '../../../utils/websocket';
import { AsashishiAgentContext } from '../../../context';
import styles from "../index.module.css";

const ShellCard: React.FC = (): JSX.Element => {
    const [shellOutput, setShellOutput] = useState<string>("");
    const { tab, ioHistories } = useContext(AsashishiAgentContext);
    const cardRef: React.Ref<HTMLDivElement> | undefined = useRef<null>(null);
    wsInstance.injectDependencies([{
        key: "shellOutput",
        value: shellOutput,
        setValue: setShellOutput,
    }]);
    useEffect(() => {
        if (cardRef.current) {
            cardRef.current?.scrollTo({
              top: cardRef.current.scrollHeight,
              behavior: "smooth",
            });
        }
        if (!ioHistories.length) {
            setShellOutput("");
        }
    }, [ioHistories, shellOutput]);
    return (
        <div
            ref={cardRef}
            className={styles.IOOutputWrapper}
            style={{ display: tab === "shell" ? "" : "none" }}
        >
            {shellOutput && <div className={styles.IOOutput}>{shellOutput}</div>}
            {ioHistories.length && ioHistories[ioHistories.length - 1].diplayPosition === "shell" && (
                <div className={styles.UserInput}>
                    <span>
                        User
                    </span>
                    {ioHistories[ioHistories.length - 1].content}
                </div>
            )}
        </div>
    );
};

export default ShellCard;
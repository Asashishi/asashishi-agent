import {
    useContext,
    useState,
} from 'react';
import type {
    JSX,
    ChangeEvent
} from 'react';
import wsInstance from '../../utils/websocket';
import styles from "./index.module.css";
import { AsashishiAgentContext } from '../../context';
import type { DisplayMsg } from '../../types/websocket_type';

const UInput: React.FC = (): JSX.Element => {
    const [uInput, setUInput] = useState<string>("");
    const { setIOHistories } = useContext(AsashishiAgentContext);
    wsInstance.injectDependencies([{
        key: "uInput",
        value: uInput,
        setValue: setUInput,
    }]);
    return (
        <div className={styles.InputBarWrapper}>
            <div className={styles.InputBar}>
                <textarea
                    name="uInput"
                    value={uInput}
                    className={styles.InputBarTextarea}
                    onChange={(event: ChangeEvent<HTMLTextAreaElement>): void => setUInput(event.target.value)}
                />
                <div className={styles.InputBarButtonsWrapper}>
                    <div
                        className={styles.InputBarButton}
                        onClick={() => setUInput("")}
                    >
                        Clear
                        <span>⌫</span>
                    </div>
                    <div
                        className={styles.InputBarButton}
                        onClick={() => {
                            setIOHistories((prev: DisplayMsg[]): DisplayMsg[] => {
                                const next: DisplayMsg[] = prev;
                                next.push({
                                    type: "input",
                                    content: uInput,
                                });
                                return next;
                            });
                            wsInstance.send({
                                type: 'user_input',
                                content: uInput,
                            });
                            setUInput("");
                        }}
                    >
                        Send
                        <span>↑</span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default UInput;
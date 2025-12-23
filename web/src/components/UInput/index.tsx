import {
    useState,
} from 'react';
import type {
    JSX,
    ChangeEvent
} from 'react';
import wsInstance from '../../utils/websocket';
import styles from "./index.module.css";

const UInput: React.FC = (): JSX.Element => {
    const [uInput, setUInput] = useState<string>("");
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
                            setUInput("");
                            wsInstance.send({
                                type: 'user_input',
                                content: uInput,
                            });
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
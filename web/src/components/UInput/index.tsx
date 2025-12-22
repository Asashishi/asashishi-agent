import {
    useState,
    Fragment
} from 'react';
import type {
    JSX,
    ChangeEvent,
} from 'react';
import wsInstance from '../../utils/websocket';

const UInput: React.FC = (): JSX.Element => {
    const [uInput, setUInput] = useState<string>("");
    wsInstance.injectDependencies([{
        key: "uInput",
        value: uInput,
        setValue: setUInput,
    }]);
    return (
        <div>
            <label htmlFor="uInput">Input: </label>
            <div>
                <textarea
                    id="uInput"
                    rows={10}
                    cols={50}
                    value={uInput}
                    onChange={(event: ChangeEvent<HTMLTextAreaElement>) =>
                        setUInput(event.target.value)
                    }
                />
            </div>
            <button onClick={() =>
                wsInstance.send({
                    type: 'user_input',
                    content: uInput,
                })
            }>
                Send
            </button>
        </div>
    );
};

export default UInput;
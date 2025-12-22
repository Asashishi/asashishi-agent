import {
    useState,
    useEffect,
    Fragment
} from 'react';
import type {
    JSX,
    ChangeEvent,
} from 'react';
import { UInputContext } from '../../context';
import wsInstance from '../../utils/websocket';
import { ContextStorageItem } from '../../utils/context_storage';

const UInput: React.FC = (): JSX.Element => {
    const uInput = new ContextStorageItem(...useState<string>(""));
    useEffect(() => {
        UInputContext.set([{
            key: "uInput",
            contextItem: uInput,
        }]);
        wsInstance.injectDependencies([{
            key: "uInput",
            contextStorage: UInputContext,
        }])
    }, []);
    return (
        <Fragment>
            <label htmlFor="uInput">Input: </label>
            <div>
                <textarea
                    id="uInput"
                    rows={10}
                    cols={50}
                    value={uInput.value}
                    onChange={(event: ChangeEvent<HTMLTextAreaElement>) =>
                        uInput.setValue(event.target.value)
                    }
                />
            </div>
            <button onClick={() =>
                wsInstance.send({
                    type: 'user_input',
                    content: uInput.value,
                })
            }>
                Send
            </button>
        </Fragment>
    );
};

export default UInput;
import React, {
    useState,
    useEffect,
    Fragment
} from 'react';
import type { JSX } from 'react';
import { ShellCardContext } from '../../context';
import wsInstance from '../../utils/websocket';
import { ContextStorageItem } from '../../utils/context_storage';

const ShellCard: React.FC = (): JSX.Element => {
    const shellOutput: ContextStorageItem<string> = new ContextStorageItem(...useState<string>(""));
    useEffect(() => {
        ShellCardContext.set([{
            key: "shellOutput",
            contextItem: shellOutput,
        }]);
        wsInstance.injectDependencies([{
            key: "shellOutput",
            contextStorage: ShellCardContext,
        }]);
    }, []);
    return (
        <Fragment>
            <label>SCP: </label>
            <button onClick={() => shellOutput.setValue("")}>Clear</button>
            <div id="scp">{shellOutput.value}</div>
        </Fragment>
    );
};

export default ShellCard;
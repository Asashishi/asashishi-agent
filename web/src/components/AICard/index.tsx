import {
    useState,
    useEffect,
    Fragment
} from 'react';
import type { JSX } from 'react';
import { AICardContext } from '../../context';
import wsInstance from '../../utils/websocket';
import { ContextStorageItem } from '../../utils/context_storage';

const AICard: React.FC = (): JSX.Element => {
    const aiOutput = new ContextStorageItem(...useState<string>(""));
    useEffect(() => {
        AICardContext.set([{
            key: "aiOutput",
            contextItem: aiOutput,
        }]);
        wsInstance.injectDependencies([{
            key: "aiOutput",
            contextStorage: AICardContext,
        }]);
    }, []);
    return (
        <Fragment>
            <label>AI: </label>
            <button onClick={() => aiOutput.setValue("")}>Clear</button>
            <div id="ai">{aiOutput.value}</div>
        </Fragment>
    );
};

export default AICard;
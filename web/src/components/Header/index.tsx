import { useContext, useState, type JSX } from 'react';
import styles from './index.module.css';
import type { SignalStaus } from '../../types/signal_status_type';
import wsInstance from '../../utils/websocket';
import { GreySignal, YellowSignal } from '../../consts/signal';
import syncServerStaus from '../../http';
import { AsashishiAgentContext } from '../../context';
import { ReGetStatusDelay } from '../../consts/http';

const Header: React.FC = (): JSX.Element => {
    const { setTab } = useContext(AsashishiAgentContext);
    const [wsSignal, setWsSignal] = useState<SignalStaus>({
        strength: 2,
        color: YellowSignal,
    });
    wsInstance.injectDependencies([{
        key: "wsSignal",
        value: wsSignal,
        setValue: setWsSignal,
    }]);
    return (
        <div className={styles.HeaderWrapper}>
            <h2>
                WebSocket Status:
            </h2>
            <p className={styles.Signal}>
                <span style={{ background: wsSignal.strength > 0 ? wsSignal.color : GreySignal }} />
                <span style={{ background: wsSignal.strength > 1 ? wsSignal.color : GreySignal }} />
                <span style={{ background: wsSignal.strength > 2 ? wsSignal.color : GreySignal }} />
                <span style={{ background: wsSignal.strength > 3 ? wsSignal.color : GreySignal }} />
            </p>
            <button
                onClick={() => {
                    wsInstance.reconnect();
                    setTimeout(() => syncServerStaus(setTab), ReGetStatusDelay);
                }}
                className={styles.ReconnectButton}
            >
                <span>⟳</span>
            </button>
        </div>
    );
}

export default Header;
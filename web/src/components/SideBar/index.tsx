import {
    useContext
} from 'react';
import type { JSX } from 'react';
import { AsashishiAgentContext } from '../../context';
import styles from "./index.module.css";
import wsInstance from '../../utils/websocket';

const SideBar: React.FC = (): JSX.Element => {
    const { tab, setTab } = useContext(AsashishiAgentContext);

    wsInstance.injectDependencies([{
        key: "tab",
        value: tab,
        setValue: setTab,
    }]);

    return(
        <div className={styles.SideBar}>
            <h2 className={styles.SideBarTitle}>
                <img src='./app.ico' className={styles.TitleIcon}/>
                Asashishi Agent
            </h2>
            <div>
                <div
                    className={styles.OptionsNewSession}
                    onClick={() => {}}
                >
                    New Session
                    <span>⟳</span>
                </div>
                <div
                    onClick={() => setTab("chat")}
                    className={
                        tab === "chat"
                        ? styles.OptionsActivite
                        : styles.Options
                    }
                >
                    AI Chat
                </div>
                <div
                    onClick={() => setTab("shell")}
                    className={
                        tab === "shell"
                        ? styles.OptionsActivite
                        : styles.Options
                    }
                > 
                    Web Shell
                </div>
            </div>
        </div>
    )
};

export default SideBar;
import {
    useContext
} from 'react';
import type { JSX } from 'react';
import { AsashishiAgentContext } from '../../context';
import styles from "./index.module.css";
import wsInstance from '../../utils/websocket';
import { RequestNewSession } from '../../consts/websocket';

const SideBar: React.FC = (): JSX.Element => {
    const { tab, setTab, setIOHistories } = useContext(AsashishiAgentContext);
    return(
        <div className={styles.SideBar}>
            <h2 className={styles.SideBarTitle}>
                <img src='./app.ico' className={styles.TitleIcon}/>
                Asashishi Agent
            </h2>
            <div>
                <div
                    className={styles.OptionsNewSession}
                    onClick={() => {
                        wsInstance.send({ type: RequestNewSession });
                        setTab("chat");
                        setIOHistories([]);
                    }}
                >
                    New Session
                    <span>⟳</span>
                </div>
                <div
                    className={
                        tab === "chat"
                        ? styles.OptionsActivite
                        : styles.Options
                    }
                >
                    AI Chat
                </div>
                <div
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
import {
    useContext
} from 'react';
import type { JSX } from 'react';
import { AsashishiAgentContext } from '../../context';
import styles from "./index.module.css";

const SideBar: React.FC = (): JSX.Element => {
    const { tab, setTab } = useContext(AsashishiAgentContext);
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
                    NewSession
                    <span>⟳</span>
                </div>
                <div
                    onClick={() => setTab("AI")}
                    className={
                        tab === "AI"
                        ? styles.OptionsActivite
                        : styles.Options
                    }
                >
                    AI Chat
                </div>
                <div
                    onClick={() => setTab("Shell")}
                    className={
                        tab === "Shell"
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
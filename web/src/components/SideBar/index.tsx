import {
    useState,
    useEffect
} from 'react';
import type { JSX } from 'react';
import styles from "./index.module.css";
import { SideBarContext } from '../../context';
import { ContextStorageItem } from '../../utils/context_storage';

type Tab = "AI" | "Shell";

const SideBar: React.FC = (): JSX.Element => {
    const tab = new ContextStorageItem(...useState<Tab>("AI"));
    useEffect(() => {
        SideBarContext.set([{
            key: "tab",
            contextItem: tab,
        }]);
    }, []);
    return(
        <div className={styles.SideBar}>
            <h2 className={styles.SideBarTitle}>
                <img src='./app.ico' className={styles.TitleIcon}/>
                Asashishi Agent
            </h2>
            <div
                onClick={() => tab.setValue("AI")}
                className={
                    tab.value === "AI"
                    ? styles.OptionsActivite
                    : styles.Options
                }
            >
                AI Chat
            </div>
            <div
                onClick={() => tab.setValue("Shell")}
                className={
                    tab.value === "Shell"
                    ? styles.OptionsActivite
                    : styles.Options
                }
            > 
                Web Shell
            </div>
        </div>
    )
};

export default SideBar;
import type {
    JSX,
} from 'react';
import type React from 'react';
import Header from '../Header';
import UInput from '../UInput';
import AICard from '../AICard';
import ShellCard from '../ShellCard';
import styles from "./index.module.css";

const Content: React.FC = (): JSX.Element => {
    return (
        <div className={styles.Content}>
            <Header />
            <UInput />
            <AICard />
            <ShellCard />
        </div>
    );
}

export default Content;
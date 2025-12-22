import type {
    JSX,
} from 'react';
import Header from './components/Header';
import SideBar from './components/SideBar';
import styles from "./app.module.css";
import UInput from './components/UInput';
import AICard from './components/AICard';
import ShellCard from './components/ShellCard';

const App: React.FC = (): JSX.Element => {
    return (
        <div className={styles.app}>
            <SideBar />
            <div>
                <Header />
                <div>
                    <UInput />
                </div>
                <AICard />
                <ShellCard />
            </div>
        </div>
    );
};

export default App;

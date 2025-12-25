import type {
    JSX,
} from 'react';
import SideBar from './components/SideBar';
import Content from './components/Content';
import styles from "./app.module.css";

const App: React.FC = (): JSX.Element => {
    return (
        <div className={styles.App}>
            <SideBar />
            <Content />
        </div>
    );
};

export default App;

import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import TextInput from '../common/TextInput';
import CustomButton from '../common/Button';
import { login } from '../../services/auth';

const LoginPage = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            await login(username, password);
        } catch (error) {
            setError(error.message);
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <TextInput
                    label="Username"
                    name="username"
                    value={username}
                    onChange={(event) => setUsername(event.target.value)}
                />
                <TextInput
                    label="Password"
                    name="password"
                    type="password"
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                />
                {error && <p>{error}</p>}
                <CustomButton type="submit">Login</CustomButton>
            </form>
            <p>
                Don't have an account? <Link to="/signup">Sign up</Link>
            </p>
        </div>
    );
};

export default LoginPage;
import React from 'react';

const TextInput = (props) => {
    const { label, name, type, value, onChange, error } = props;

    return (
        <div className="form-group">
            <label htmlFor={name}>{label}</label>
            <input
                type={type}
                name={name}
                value={value}
                onChange={onChange}
                className={`form-control ${error ? 'is-invalid' : ''}`}
            />
            {error && <div className="invalid-feedback">{error}</div>}
        </div>
    );
};

export default TextInput;
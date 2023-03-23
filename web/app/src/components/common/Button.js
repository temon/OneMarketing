import React from 'react';
import { Button } from 'antd';

const CustomButton = ({ children, ...props }) => (
    <Button {...props}>
        {children}
    </Button>
);

export default CustomButton;
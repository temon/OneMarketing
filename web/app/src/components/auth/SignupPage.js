import React, { useState } from 'react';
import { Form, Input, Checkbox } from 'antd';
import CustomButton from "../common/Button";

const SignupPage = () => {
    const [form] = Form.useForm();
    const [loading, setLoading] = useState(false);

    const onFinish = async (values) => {
        setLoading(true);
        try {
            // call your API or authentication service here to create new user
            console.log(values);
        } catch (error) {
            console.log('Signup failed:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="signup-container">
            <h1>Signup Page</h1>
            <Form form={form} onFinish={onFinish} className="signup-form">
                <Form.Item
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input placeholder="Username" />
                </Form.Item>

                <Form.Item
                    name="email"
                    rules={[{ required: true, message: 'Please input your email!' }]}
                >
                    <Input placeholder="Email" />
                </Form.Item>

                <Form.Item
                    name="password"
                    rules={[{ required: true, message: 'Please input your password!' }]}
                >
                    <Input.Password placeholder="Password" />
                </Form.Item>

                <Form.Item
                    name="confirm"
                    dependencies={['password']}
                    rules={[
                        {
                            required: true,
                            message: 'Please confirm your password!',
                        },
                        ({ getFieldValue }) => ({
                            validator(_, value) {
                                if (!value || getFieldValue('password') === value) {
                                    return Promise.resolve();
                                }
                                return Promise.reject(
                                    new Error('The two passwords that you entered do not match!')
                                );
                            },
                        }),
                    ]}
                >
                    <Input.Password placeholder="Confirm Password" />
                </Form.Item>

                <Form.Item name="remember" valuePropName="checked">
                    <Checkbox>Remember me</Checkbox>
                </Form.Item>

                <Form.Item>
                    <CustomButton type="primary" htmlType="submit" loading={loading}>
                        Signup
                    </CustomButton>
                </Form.Item>
            </Form>
        </div>
    );
};

export default SignupPage;
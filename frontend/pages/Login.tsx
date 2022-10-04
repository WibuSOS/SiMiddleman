import { Form } from "react-bootstrap";

export default function LoginForm({ handleSubmit }) {
    return (
        <Form onSubmit={handleSubmit} id="loginForm">
            <Form.Group className="mb-3" controlId='userEmail'>
                <Form.Control
                type="email"
                placeholder="Email"
                name="email"
                required
                autoFocus/>
            </Form.Group>
            <Form.Group className="mb-3" controlId='userPassword'>
                <Form.Control
                type="password"
                placeholder="Password"
                name="password"
                autoFocus/>
            </Form.Group>
        </Form>
    )
}
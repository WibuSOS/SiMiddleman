import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

export default function RegisterForm({ handleSubmit }) {
    return (        
        <Form onSubmit={handleSubmit}>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                <Form.Control
                    type="text"
                    placeholder="Nama"
                    name='nama'
                    autoFocus
                    required
                />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput2">
                <Form.Control
                    type="text"
                    placeholder="No HP"
                    name='noHp'
                    required
                />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput3">
                <Form.Control
                    type="text"
                    placeholder="No Rekening"
                    name='noRek'
                    required
                />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput4">
                <Form.Control
                    type="email"
                    placeholder="Email"
                    name='email'
                    required
                />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput5">
                <Form.Control
                    type="password"
                    placeholder="Password"
                    name='password'
                    required
                    minLength={8}
                />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput6">
                <Form.Control
                    type="password"
                    placeholder="Confirm Password"
                    name='confirmPassword'
                    required
                    minLength={8}
                />
            </Form.Group>
            <Button variant='merah' type='submit'>Submit</Button>
        </Form>
    )
}
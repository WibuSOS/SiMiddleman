import { Form, Button } from 'react-bootstrap';

export default function FormRegister({ handleSubmitRegister }) {
    return (
        <Form onSubmit={handleSubmitRegister}>
            <Form.Group className="mb-3">
              <Form.Control
                  type="text"
                  placeholder="Nama"
                  name='nama'
                  data-testid="nama"
                  autoFocus
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="text"
                  placeholder="No HP"
                  name='noHp'
                  data-testid="noHp"
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="text"
                  placeholder="No Rekening"
                  name='noRek'
                  data-testid="noRek"
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="email"
                  placeholder="Email"
                  name='email'
                  data-testid="email"
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="password"
                  placeholder="Password"
                  name='password'
                  data-testid="password"
                  required
                  minLength={8}
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="password"
                  placeholder="Confirm Password"
                  name='confirmPassword'
                  data-testid="confirmPassword"
                  required
                  minLength={8}
              />
            </Form.Group>
            <Button variant='merah' type='submit' className='w-100'>Register</Button>
        </Form>
    )
}
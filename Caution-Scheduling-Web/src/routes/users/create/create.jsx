import { useEffect } from "react";
import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import "../users.css"

function NewUserButton() {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const [userName, setUserName] = useState('')
    const [firstName, setFirstName] = useState('')
    const [lastName, setLastName] = useState('')
    const [Password, setPassword] = useState('')
    const [Password2, setPassword2] = useState('')
    const [email, setEmail] = useState('')
    const [role, setRole] = useState('')

    

    const createNewUser = async (userName, firstName, lastName, Password, email, role) => {
        const myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
      
        const urlencoded = new URLSearchParams();
        urlencoded.append('userName', username);
        urlencoded.append('firstName', firstName);
        urlencoded.append('lastName', lastName);
        urlencoded.append('password', Password);
        urlencoded.append('email', email);
        urlencoded.append('role', role);
      
        const requestOptions = {
          method: 'POST',
          headers: myHeaders,
          body: urlencoded,
          redirect: 'follow',
        };
      
        try {
          const response = await fetch('/api/luser/login', requestOptions);
      
          if (!response.ok) {
            throw new Error('Invalid login credentials.');
          }
      
          const data = await response.json();
          return data;
        } catch (error) {
          console.error('Login failed:', error);
          throw new Error('Login failed. Please try again.');
        }
      };

    return (
        <>
            <button
                className="add-new-user-button"
                onClick={handleShow}
            >
                Add new user
            </button>

            <Modal show={show} onHide={handleClose} backdrop="static">
                <Modal.Header closeButton>
                    <Modal.Title>Create a New User</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>UserName</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={userName}
                            onChange={(e) => setUserName(e.target.value)}
                            placeholder="UserName"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Password</strong>
                        </Form.Label>
                        <Form.Control
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="Password"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Confirm Password</strong>
                        </Form.Label>
                        <Form.Control
                            type="password"
                            value={password2}
                            onChange={(e) => setPassword2(e.target.value)}
                            placeholder="Password"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>First Name</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={firstName}
                            onChange={(e) => setFirstName(e.target.value)}
                            placeholder="First Name"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Last Name</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={lastName}
                            onChange={(e) => setLastName(e.target.value)}
                            placeholder="Last Name"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Email</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="Email"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Role</strong>
                        </Form.Label>
                        <Form.Select aria-label="Default select example" onChange={(e) => setRole(e.target.value)}>
                            <option value="1">Student</option>
                            <option value="2">Tutor</option>
                            <option value="3">Supervisor</option>
                            <option value="4">Administrator</option>
                        </Form.Select>
                    </Form.Group>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="danger" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="success" className="background-1" onClick={function () { handleClose; toggle() }}>
                        Create new user
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )

}


export default NewUserButton
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
    const [password, setPassword] = useState('')
    const [email, setEmail] = useState('')
    const [role, setRole] = useState('')
    const [userCreate, setUserCreate] = useState(false)

    

    const createNewUser = async () => {
        const myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
      
        const urlencoded = new URLSearchParams();
        urlencoded.append('userName', userName);
        urlencoded.append('firstName', firstName);
        urlencoded.append('lastName', lastName);
        urlencoded.append('password', password);
        urlencoded.append('email', email);
        urlencoded.append('role', role);
      
        const requestOptions = {
          method: 'POST',
          headers: myHeaders,
          body: urlencoded,
          redirect: 'follow',
        };
      
        try {
          const response = await fetch('/api/luser', requestOptions);
      
          if (!response.ok) {
            throw new Error('Something Went Wrong');
          }
      
          const data = await response.json();
          setUserCreate(true)
          //return data;

        } catch (error) {
            alert(error)
          console.error('Error:', error);
          throw new Error('Error');
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
                            <option value="0">Select a Role</option>
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
                    <Button variant="success" className="background-1" onClick={function () {createNewUser();  handleClose(); }}>
                        Create new user
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )

}


export default NewUserButton
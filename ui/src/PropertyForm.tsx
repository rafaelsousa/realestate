import React from "react";
import {Button, Container, Form, Jumbotron} from "react-bootstrap";

class PropertyForm extends React.Component {

    state: {
        value: any;
    }

    constructor(props: any) {
        super(props);
        this.state = {
            value: 'Please write an essay about your favorite DOM element.'
        };
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit(event: any) {
        alert('the form was submitted');
        event.preventDefault();
    }

    render() {
        return (
            <Container fluid className={"card-body"}>
                <Jumbotron>
                    <h1>Realestate Chain</h1>
                    <p>The home of the Domus token</p>
                </Jumbotron>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Group controlId="formBasicAddress">
                        <Form.Label>Address</Form.Label>
                        <Form.Control type="text" placeholder="Enter property address" />
                    </Form.Group>
                    <Form.Group controlId="formCity">
                        <Form.Label>City</Form.Label>
                        <Form.Control type="text" placeholder="Enter property city" />
                    </Form.Group>
                    <Form.Group controlId="formState">
                        <Form.Label>Province/State</Form.Label>
                        <Form.Control type="text" placeholder="Enter property province / state" />
                    </Form.Group>
                    <Form.Group controlId="formCountry">
                        <Form.Label>Country</Form.Label>
                        <Form.Control type="text" placeholder="Enter property Country" />
                    </Form.Group>
                    <Button type="submit">
                        Submit
                    </Button>
                </Form>
            </Container>
        );
    }
}

export default PropertyForm;

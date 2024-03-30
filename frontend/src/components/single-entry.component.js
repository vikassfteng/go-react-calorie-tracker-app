
import React from 'react';

import 'bootstrap/dist/css/bootstrap.css'

import {Button, Card, Row, Col} from 'react-bootstrap'

const Entry =({entryData, setChangeIngredient, deleteSingleEntry, setChangeEntry}) => {
    return(
        <Card>
            <Row>
                <Col>Dish:{entryData !== undefined && entryData.dish}</Col>
                <Col>Ingredients:{entryData !== undefined && entryData.ingredients}</Col>
                <Col>Calories:{entryData !== undefined && entryData.calories}</Col>
                <Col>Fat:{entryData !== undefined && entryData.fat}</Col>
                <Col><Button variant="danger" onClick={()=> deleteSingleEntry(entryData._id)}>
                <i className="fas fa-trash-alt"></i> delete entry</Button></Col>
                <Col><Button variant="info" onClick={()=> changeIngredient()}>
                <i className="fas fa-pencil-alt"></i>  change ingredients</Button></Col>
                <Col><Button variant="warning" onClick={()=> changeEntry()}>
                <i className="fas fa-edit"></i>change entry</Button></Col>
            </Row>
        </Card>
    )

    function changeIngredient(){
        setChangeIngredient(
            {
                "change": true,
                "id":entryData._id
            }
        )
    }

    function changeEntry(){
        setChangeEntry(
            {
                "change": true,
                "id":entryData._id
            }
        )
    }
}

export default Entry;

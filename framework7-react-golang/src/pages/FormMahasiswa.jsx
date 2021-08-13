import React, { Component, useEffect, useState } from "react";
import WithListLoading from "../components/withListLoading";
import axios from "axios";
import {
  Page,
  Navbar,
  List,
  ListInput,
  ListItem,
  Toggle,
  BlockTitle,
  Row,
  Button,
  Range,
  Block,
  Col,
} from "framework7-react";

const endpoint = `http://localhost:8080/mahasiswa`;

const FormMahasiswa = () => {
  const [appState, setAppState] = useState("");

  useEffect(() => {
    const apiUrl = `http://localhost:8080/mahasiswa`;
    fetch(apiUrl)
      .then((res) => res.json())
      .then((mahasiswa) => {
        setAppState({ mahasiswa: mahasiswa });
      });
  }, [setAppState]);
  console.log(appState);

  return (
    <>
      <Page name="form">
        <Navbar title="Form Tambah Mahasiswa" backLink="Back"></Navbar>

        <BlockTitle>Tambah Mahasiswa</BlockTitle>
        <List noHairlinesMd>
          <ListInput label="Nim" type="text" placeholder="NIM"></ListInput>

          <ListInput label="Nama" type="text" placeholder="Nama"></ListInput>

          <Row style={{ marginTop: "30px" }}>
            <Col />
            <Col>
              <Button fill>Submit</Button>
            </Col>
            <Col />
          </Row>
        </List>
      </Page>
    </>
  );
};

export default FormMahasiswa;

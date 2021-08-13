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

let endpoint = "http://localhost:8080";

const FormPage = () => {
  const [nim, setNim] = useState("");
  const [nama, setNama] = useState("");
  const [items, setItems] = useState([]);

  const onChangeNim = (event) => {
    console.log(event.target.name);
    setNim(event.target.value);
  };

  const onChangeNama = (event) => {
    console.log(event.target.name);
    setNama(event.target.value);
  };

  const onSubmit = async () => {
    if (nim && nama) {
      await axios
        .post(
          endpoint + "/mahasiswa",
          {
            nim,
            nama,
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          }
        )
        .then((res) => {
          console.log(res);
        });
    }
  };

  return (
    <>
      <Page name="form">
        <Navbar title="Form Tambah Mahasiswa" backLink="Back"></Navbar>

        <BlockTitle>Tambah Mahasiswa</BlockTitle>
        <form action="" onSubmit={onSubmit}>
          <List noHairlinesMd>
            <ListInput
              label="Nim"
              type="text"
              placeholder="NIM"
              name="nim"
              onChange={onChangeNim}
              value={nim}
            ></ListInput>

            <ListInput
              label="Nama"
              type="text"
              placeholder="Nama"
              name="nama"
              onChange={onChangeNama}
              value={nama}
            ></ListInput>

            <Row style={{ marginTop: "30px" }}>
              <Col />
              <Col>
                <Button type="submit" fill>
                  Submit
                </Button>
              </Col>
              <Col />
            </Row>
          </List>
        </form>
      </Page>
    </>
  );
};

export default FormPage;

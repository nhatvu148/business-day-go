import React, { useState } from "react";
import {
  Button,
  Grid,
  List,
  Typography,
  IconButton,
  styled,
  Box,
  Card,
} from "@mui/material";
import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import ListItem from "@mui/material/ListItem";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import ListItemText from "@mui/material/ListItemText";
import Avatar from "@mui/material/Avatar";
import AddIcon from "@mui/icons-material/Add";
import DeleteIcon from "@mui/icons-material/Delete";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";

const Demo = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
}));

const initialCustomHolidays = [
  "2022/02/02",
  "2022/03/03",
  "2022/04/04",
  "2022/05/05",
];

function generate(element: React.ReactElement) {
  return [0, 1, 2].map((value) =>
    React.cloneElement(element, {
      key: value,
    })
  );
}

const Home: NextPage = () => {
  const [customHolidays, setCustomHolidays] = useState(initialCustomHolidays);
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        position: "relative",
        marginTop: 15,
      }}
    >
      <Box sx={{ width: "40%" }}>
        <Card variant="outlined" style={{ padding: 10 }}>
          <div style={{ display: "flex", justifyContent: "center" }}>
            <Typography
              sx={{ mt: 1, mb: 1, mr: 2 }}
              variant="h6"
              component="div"
            >
              Custom Holidays
            </Typography>
            <IconButton
              edge="end"
              aria-label="add"
              onClick={() => {
                setCustomHolidays((prev) => [...prev, "2022/05/06"]);
              }}
            >
              <AddIcon />
            </IconButton>
          </div>

          <Demo>
            <List dense={false}>
              {customHolidays.map((day: string, id: number) => {
                return (
                  <ListItem
                    key={id}
                    secondaryAction={
                      <IconButton edge="end" aria-label="delete">
                        <DeleteIcon />
                      </IconButton>
                    }
                  >
                    <ListItemAvatar>
                      <Avatar>
                        <CalendarMonthIcon />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText
                      primary={day}
                      secondary={true ? "Secondary text" : null}
                    />
                  </ListItem>
                );
              })}
            </List>
          </Demo>
        </Card>
      </Box>
    </div>
  );
};

export default Home;

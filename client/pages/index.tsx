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
  Tooltip
} from "@mui/material";
import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import ListItem from "@mui/material/ListItem";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import ListItemText from "@mui/material/ListItemText";
import Avatar from "@mui/material/Avatar";
import AddBoxIcon from "@mui/icons-material/AddBox";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";
import { pink } from "@mui/material/colors";

const Demo = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
}));

const initialCustomHolidays = [
  { date: "2022/02/02", category: "Holiday" },
  { date: "2022/03/03", category: "Business day" },
  { date: "2022/04/04", category: "Holiday" },
  { date: "2022/05/05", category: "Business day" },
  { date: "2022/06/06", category: "Holiday" },
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
            <Tooltip title="Add Custom Holiday">
              <IconButton
                edge="end"
                aria-label="add"
                onClick={() => {
                  setCustomHolidays((prev) => [
                    ...prev,
                    { date: "2022/05/06", category: "Holiday" },
                  ]);
                }}
              >
                <AddBoxIcon color="primary" />
              </IconButton>
            </Tooltip>
          </div>

          <Demo>
            <List dense={false}>
              {customHolidays.map((day, id: number) => {
                return (
                  <ListItem
                    key={id}
                    secondaryAction={
                      <>
                        <Tooltip title="Edit">
                          <IconButton
                            edge="end"
                            aria-label="edit"
                            sx={{ mr: 0.1 }}
                          >
                            <EditIcon color="success" />
                          </IconButton>
                        </Tooltip>
                        <Tooltip title="Delete">
                          <IconButton edge="end" aria-label="delete">
                            <DeleteIcon sx={{ color: pink[500] }} />
                          </IconButton>
                        </Tooltip>
                      </>
                    }
                  >
                    <ListItemAvatar>
                      <Avatar>
                        <CalendarMonthIcon color="primary" />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText primary={day.date} secondary={day.category} />
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

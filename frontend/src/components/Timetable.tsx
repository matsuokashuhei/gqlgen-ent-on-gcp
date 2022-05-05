import { TabContext, TabList, TabPanel } from "@mui/lab";
import { Box, Grid, Tab, Tabs } from "@mui/material";
import { formatRFC3339, parse } from "date-fns";
import { useEffect, useState, VFC, SyntheticEvent } from "react";
import { Link } from "react-router-dom";
import { PartialDeep } from "type-fest";
import {
  GetClassesBySchoolQuery,
  useGetClassesBySchoolLazyQuery,
} from "../generated/graphql";

type Props = {
  date: string;
  onClickClass: (clazz: ClassType) => void;
  onClickNewClass?: (schedule: PartialDeep<SchedulesType[number]>) => void;
};

type StudiosType = GetClassesBySchoolQuery["school"]["studios"];
type StudioType = StudiosType[number];
type RoomsType = StudioType["rooms"];
type RoomType = RoomsType[number];
type SchedulesType = RoomType["schedules"];
type ClassType = NonNullable<SchedulesType[number]["class"]>;
type InstructorType = ClassType["instructor"];

export const Timetable: VFC<Props> = ({ date, onClickClass }) => {
  const [getClassesBySchool, { data, loading, error }] =
    useGetClassesBySchoolLazyQuery();

  const [studios, setStudios] = useState<StudiosType>([]);
  const [studioId, setStudioId] = useState<string | null>(null);
  const [roomId, setRoomId] = useState<string | null>(null);

  useEffect(() => {
    getClassesBySchool({
      variables: {
        id: "1",
        date: formatRFC3339(parse(date, "yyyyMMdd", new Date())),
      },
    });
  }, [date, getClassesBySchool]);

  useEffect(() => {
    if (!data) return;
    setStudios(data.school.studios);
  }, [data]);

  useEffect(() => {
    if (!studios || studios.length === 0) return;

    setStudioId(studios[0].id);
  }, [studios]);

  useEffect(() => {
    if (!studioId) return;

    // setRoomId(studios.find((studio) => studio.id === studioId).rooms[0].id);
  }, [studioId]);

  const handleChangeStudio = (event: SyntheticEvent, newValue: string) => {
    setStudioId(newValue);
  };

  const handleChangeRoom = (event: SyntheticEvent, newValue: string) => {
    setRoomId(newValue);
  };

  const renderTabGroupForStudios = (studios: StudiosType) => {
    if (!studios || !studioId) return;
    return (
      <Box>
        <TabContext value={studioId}>
          <Box>
            <TabList onChange={handleChangeStudio}>
              {studios.map((studio) => (
                <Tab
                  key={studio.id}
                  label={studio.name}
                  value={studio.id}
                ></Tab>
              ))}
            </TabList>
          </Box>
          {studios.map((studio) => (
            <TabPanel key={studio.id} value={studio.id}>
              {renderTabGroupForRooms(studio.rooms)}
            </TabPanel>
          ))}
        </TabContext>
      </Box>
    );
  };

  const renderTabGroupForRooms = (rooms: RoomsType) => {
    if (!rooms || !roomId) return;
    return (
      <Box>
        <TabContext value={roomId}>
          <Box>
            <TabList onChange={handleChangeRoom}>
              {rooms.map((room) => (
                <Tab key={room.id} label={room.name} value={room.id} />
              ))}
            </TabList>
          </Box>
          {rooms.map((room) => (
            <TabPanel key={room.id} value={room.id}>
              {renderSchedules(room.schedules)}
            </TabPanel>
          ))}
        </TabContext>
      </Box>
    );
  };

  const renderSchedules = (schedules: SchedulesType) => {
    return (
      <Grid container columns={8}>
        {[0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
          return (
            <Grid item xs={1} key={dayOfWeek}>
              {dayOfWeek}
            </Grid>
          );
        })}
        {renderTimeslots(schedules)}
      </Grid>
    );
  };

  const renderTimeslots = (schedules: SchedulesType) => {
    return [
      { startTime: "13:00", endTime: "14:00" },
      { startTime: "14:15", endTime: "15:25" },
      { startTime: "15:30", endTime: "16:40" },
      { startTime: "16:45", endTime: "17:55" },
      { startTime: "18:00", endTime: "19:10" },
      { startTime: "19:15", endTime: "20:25" },
      { startTime: "20:30", endTime: "21:40" },
      { startTime: "21:45", endTime: "22:55" },
    ].map((slot) => {
      return [0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
        if (dayOfWeek === 0) {
          return (
            <Grid item xs={1} key={dayOfWeek}>
              <div>{slot.startTime}</div>
              <div>~ {slot.endTime}</div>
            </Grid>
          );
        }
        const schedule = schedules.find(
          (schedule) =>
            schedule.dayOfWeek === dayOfWeek &&
            schedule.startTime === slot.startTime
        );
        if (!schedule)
          return <Grid item xs={1} key={`${dayOfWeek}-none`}></Grid>;
        if (schedule.class) {
          return (
            <Grid item xs={1} key={`${dayOfWeek}-${schedule.id}`}>
              {renderClass(schedule.class)}
            </Grid>
          );
        } else {
          return (
            <Grid item xs={1} key={`${dayOfWeek}-${schedule.id}`}>
              {renderNewClassButton(schedule)}
            </Grid>
          );
        }
      });
    });
  };

  const renderClass = (clazz: ClassType) => {
    if (!clazz) return <></>;
    return (
      <div onClick={() => onClickClass(clazz)}>
        <div>{clazz.name}</div>
        <div>{clazz.level}</div>
        {renderInstructor(clazz.instructor)}
      </div>
    );
  };

  const renderInstructor = (instructor: InstructorType) => {
    return <div className="text-md text-gray-700">{instructor.name}</div>;
  };

  const renderNewClassButton = (
    schedule?: PartialDeep<SchedulesType[number]>
  ) => {
    if (!schedule) return <></>;
    return <Link to={`/schedules/${schedule.id}/classes/new`}>+</Link>;
  };

  if (!data) {
    return <></>;
  }
  // const {
  //   school: { studios },
  // } = data;

  return (
    <>
      <h1>クラス</h1>
      <div className="bg-white">{renderTabGroupForStudios(studios)}</div>
    </>
  );
};

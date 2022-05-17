import { TabContext, TabList, TabPanel } from "@mui/lab";
import { Grid, Tab } from "@mui/material";
import { SyntheticEvent, VFC } from "react";
import { Link } from "react-router-dom";
import { PartialDeep } from "type-fest";
import { GetClassesBySchoolQuery } from "../generated/graphql";

type Props = {
  date: string;
  school: GetClassesBySchoolQuery["school"];
  selectedStudioId: string;
  selectedRoomId: string;
  handleSelectStudioId: (studioId: string) => void;
  handleSelectRoomId: (roomId: string) => void;
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

export const Timetable: VFC<Props> = ({
  school,
  selectedStudioId,
  selectedRoomId,
  handleSelectStudioId,
  handleSelectRoomId,
  ...props
}) => {
  const { studios } = school;
  const handleChangeStudio = (event: SyntheticEvent, newValue: string) => {
    handleSelectStudioId(newValue);
  };
  const handleChangeRoom = (event: SyntheticEvent, newValue: string) => {
    handleSelectRoomId(newValue);
  };

  const renderScheduleHeader = () =>
    [0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => (
      <Grid item xs={1} key={dayOfWeek}>
        {dayOfWeek}
      </Grid>
    ));

  const renderSchedules = (schedules: SchedulesType) =>
    [
      { startTime: "13:00", endTime: "14:00" },
      { startTime: "14:15", endTime: "15:25" },
      { startTime: "15:30", endTime: "16:40" },
      { startTime: "16:45", endTime: "17:55" },
      { startTime: "18:00", endTime: "19:10" },
      { startTime: "19:15", endTime: "20:25" },
      { startTime: "20:30", endTime: "21:40" },
      { startTime: "21:45", endTime: "22:55" },
    ].map((slot) =>
      [0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
        if (dayOfWeek === 0) {
          return (
            <Grid item xs={1} key={dayOfWeek}>
              <div>{slot.startTime}</div>
              <div>{slot.endTime}</div>
            </Grid>
          );
        } else {
          const schedule = schedules.find(
            (schedule) =>
              schedule.dayOfWeek === dayOfWeek &&
              schedule.startTime === slot.startTime
          );
          console.log("schedule", schedule);
          if (schedule) {
            if (schedule.class) {
              return (
                <Grid item xs={1} key={schedule.class.id}>
                  {schedule.class.name}
                  {schedule.class.level}
                  {schedule.class.instructor.name}
                </Grid>
              );
            } else {
              return (
                <Grid item xs={1} key={`${dayOfWeek}-${slot.startTime}`}>
                  <Link to={`/schedules/${schedule.id}/classes/new`}>+</Link>
                </Grid>
              );
            }
          } else {
            return (
              <Grid item xs={1} key={`${dayOfWeek}-${slot.startTime}`}></Grid>
            );
          }
        }
      })
    );
  return (
    <TabContext value={selectedStudioId}>
      <TabList onChange={handleChangeStudio}>
        {studios.map((studio) => (
          <Tab key={studio.id} label={studio.name} value={studio.id} />
        ))}
      </TabList>
      {studios.map((studio) => (
        <TabPanel key={studio.id} value={studio.id} sx={{ px: 0 }}>
          <TabContext value={selectedRoomId}>
            <TabList onChange={handleChangeRoom}>
              {studio.rooms.map((room) => (
                <Tab key={room.id} label={room.name} value={room.id} />
              ))}
            </TabList>
            {studio.rooms.map((room) => (
              <TabPanel key={room.id} value={room.id}>
                <Grid container columns={8}>
                  {renderScheduleHeader()}
                  {renderSchedules(room.schedules)}
                </Grid>
              </TabPanel>
            ))}
          </TabContext>
        </TabPanel>
      ))}
    </TabContext>
  );
};

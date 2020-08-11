
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TLabeledEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewLabeledEdit(owner IComponent) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = LabeledEdit_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsLabeledEdit(obj interface{}) *TLabeledEdit {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TLabeledEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsLabeledEdit.
func LabeledEditFromInst(inst uintptr) *TLabeledEdit {
    return AsLabeledEdit(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsLabeledEdit.
func LabeledEditFromObj(obj IObject) *TLabeledEdit {
    return AsLabeledEdit(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsLabeledEdit.
func LabeledEditFromUnsafePointer(ptr unsafe.Pointer) *TLabeledEdit {
    return AsLabeledEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (l *TLabeledEdit) Free() {
    if l.instance != 0 {
        LabeledEdit_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TLabeledEdit) Instance() uintptr {
    return l.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TLabeledEdit) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TLabeledEdit) IsValid() bool {
    return l.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TLabeledEdit) Is() TIs {
    return TIs(l.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TLabeledEdit) As() TAs {
//    return TAs(l.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TLabeledEditClass() TClass {
    return LabeledEdit_StaticClassType()
}

// 设置组件边界。
//
// Set component boundaries.
func (l *TLabeledEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LabeledEdit_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// 清除。
func (l *TLabeledEdit) Clear() {
    LabeledEdit_Clear(l.instance)
}

// 清除选择。
func (l *TLabeledEdit) ClearSelection() {
    LabeledEdit_ClearSelection(l.instance)
}

// 复制到粘贴板。
func (l *TLabeledEdit) CopyToClipboard() {
    LabeledEdit_CopyToClipboard(l.instance)
}

// 剪切到粘贴板。
func (l *TLabeledEdit) CutToClipboard() {
    LabeledEdit_CutToClipboard(l.instance)
}

// 从剪切板粘贴。
func (l *TLabeledEdit) PasteFromClipboard() {
    LabeledEdit_PasteFromClipboard(l.instance)
}

// 撤销上一次操作。
func (l *TLabeledEdit) Undo() {
    LabeledEdit_Undo(l.instance)
}

// 全选。
func (l *TLabeledEdit) SelectAll() {
    LabeledEdit_SelectAll(l.instance)
}

// 是否可以获得焦点。
func (l *TLabeledEdit) CanFocus() bool {
    return LabeledEdit_CanFocus(l.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (l *TLabeledEdit) ContainsControl(Control IControl) bool {
    return LabeledEdit_ContainsControl(l.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (l *TLabeledEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(LabeledEdit_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (l *TLabeledEdit) DisableAlign() {
    LabeledEdit_DisableAlign(l.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (l *TLabeledEdit) EnableAlign() {
    LabeledEdit_EnableAlign(l.instance)
}

// 查找子控件。
//
// Find sub controls.
func (l *TLabeledEdit) FindChildControl(ControlName string) *TControl {
    return AsControl(LabeledEdit_FindChildControl(l.instance, ControlName))
}

func (l *TLabeledEdit) FlipChildren(AllLevels bool) {
    LabeledEdit_FlipChildren(l.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (l *TLabeledEdit) Focused() bool {
    return LabeledEdit_Focused(l.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (l *TLabeledEdit) HandleAllocated() bool {
    return LabeledEdit_HandleAllocated(l.instance)
}

// 插入一个控件。
//
// Insert a control.
func (l *TLabeledEdit) InsertControl(AControl IControl) {
    LabeledEdit_InsertControl(l.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (l *TLabeledEdit) Invalidate() {
    LabeledEdit_Invalidate(l.instance)
}

// 移除一个控件。
//
// Remove a control.
func (l *TLabeledEdit) RemoveControl(AControl IControl) {
    LabeledEdit_RemoveControl(l.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (l *TLabeledEdit) Realign() {
    LabeledEdit_Realign(l.instance)
}

// 重绘。
//
// Repaint.
func (l *TLabeledEdit) Repaint() {
    LabeledEdit_Repaint(l.instance)
}

// 按比例缩放。
//
// Scale by.
func (l *TLabeledEdit) ScaleBy(M int32, D int32) {
    LabeledEdit_ScaleBy(l.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (l *TLabeledEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    LabeledEdit_ScrollBy(l.instance, DeltaX , DeltaY)
}

// 设置控件焦点。
//
// Set control focus.
func (l *TLabeledEdit) SetFocus() {
    LabeledEdit_SetFocus(l.instance)
}

// 控件更新。
//
// Update.
func (l *TLabeledEdit) Update() {
    LabeledEdit_Update(l.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (l *TLabeledEdit) BringToFront() {
    LabeledEdit_BringToFront(l.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (l *TLabeledEdit) ClientToScreen(Point TPoint) TPoint {
    return LabeledEdit_ClientToScreen(l.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (l *TLabeledEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (l *TLabeledEdit) Dragging() bool {
    return LabeledEdit_Dragging(l.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (l *TLabeledEdit) HasParent() bool {
    return LabeledEdit_HasParent(l.instance)
}

// 隐藏控件。
//
// Hidden control.
func (l *TLabeledEdit) Hide() {
    LabeledEdit_Hide(l.instance)
}

// 发送一个消息。
//
// Send a message.
func (l *TLabeledEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LabeledEdit_Perform(l.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (l *TLabeledEdit) Refresh() {
    LabeledEdit_Refresh(l.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (l *TLabeledEdit) ScreenToClient(Point TPoint) TPoint {
    return LabeledEdit_ScreenToClient(l.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (l *TLabeledEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (l *TLabeledEdit) SendToBack() {
    LabeledEdit_SendToBack(l.instance)
}

// 显示控件。
//
// Show control.
func (l *TLabeledEdit) Show() {
    LabeledEdit_Show(l.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (l *TLabeledEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return LabeledEdit_GetTextBuf(l.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (l *TLabeledEdit) GetTextLen() int32 {
    return LabeledEdit_GetTextLen(l.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (l *TLabeledEdit) SetTextBuf(Buffer string) {
    LabeledEdit_SetTextBuf(l.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (l *TLabeledEdit) FindComponent(AName string) *TComponent {
    return AsComponent(LabeledEdit_FindComponent(l.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (l *TLabeledEdit) GetNamePath() string {
    return LabeledEdit_GetNamePath(l.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TLabeledEdit) Assign(Source IObject) {
    LabeledEdit_Assign(l.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (l *TLabeledEdit) ClassType() TClass {
    return LabeledEdit_ClassType(l.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TLabeledEdit) ClassName() string {
    return LabeledEdit_ClassName(l.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TLabeledEdit) InstanceSize() int32 {
    return LabeledEdit_InstanceSize(l.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TLabeledEdit) InheritsFrom(AClass TClass) bool {
    return LabeledEdit_InheritsFrom(l.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (l *TLabeledEdit) Equals(Obj IObject) bool {
    return LabeledEdit_Equals(l.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TLabeledEdit) GetHashCode() int32 {
    return LabeledEdit_GetHashCode(l.instance)
}

// 文本类信息。
//
// Text information.
func (l *TLabeledEdit) ToString() string {
    return LabeledEdit_ToString(l.instance)
}

func (l *TLabeledEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    LabeledEdit_AnchorToNeighbour(l.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (l *TLabeledEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    LabeledEdit_AnchorParallel(l.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (l *TLabeledEdit) AnchorHorizontalCenterTo(ASibling IControl) {
    LabeledEdit_AnchorHorizontalCenterTo(l.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (l *TLabeledEdit) AnchorVerticalCenterTo(ASibling IControl) {
    LabeledEdit_AnchorVerticalCenterTo(l.instance, CheckPtr(ASibling))
}

func (l *TLabeledEdit) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    LabeledEdit_AnchorSame(l.instance, ASide , CheckPtr(ASibling))
}

func (l *TLabeledEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    LabeledEdit_AnchorAsAlign(l.instance, ATheAlign , ASpace)
}

func (l *TLabeledEdit) AnchorClient(ASpace int32) {
    LabeledEdit_AnchorClient(l.instance, ASpace)
}

// 获取文字对齐。
//
// Get Text alignment.
func (l *TLabeledEdit) Alignment() TAlignment {
    return LabeledEdit_GetAlignment(l.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (l *TLabeledEdit) SetAlignment(value TAlignment) {
    LabeledEdit_SetAlignment(l.instance, value)
}

// 获取四个角位置的锚点。
func (l *TLabeledEdit) Anchors() TAnchors {
    return LabeledEdit_GetAnchors(l.instance)
}

// 设置四个角位置的锚点。
func (l *TLabeledEdit) SetAnchors(value TAnchors) {
    LabeledEdit_SetAnchors(l.instance, value)
}

// 获取自动选择。
func (l *TLabeledEdit) AutoSelect() bool {
    return LabeledEdit_GetAutoSelect(l.instance)
}

// 设置自动选择。
func (l *TLabeledEdit) SetAutoSelect(value bool) {
    LabeledEdit_SetAutoSelect(l.instance, value)
}

// 获取自动调整大小。
func (l *TLabeledEdit) AutoSize() bool {
    return LabeledEdit_GetAutoSize(l.instance)
}

// 设置自动调整大小。
func (l *TLabeledEdit) SetAutoSize(value bool) {
    LabeledEdit_SetAutoSize(l.instance, value)
}

func (l *TLabeledEdit) BiDiMode() TBiDiMode {
    return LabeledEdit_GetBiDiMode(l.instance)
}

func (l *TLabeledEdit) SetBiDiMode(value TBiDiMode) {
    LabeledEdit_SetBiDiMode(l.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (l *TLabeledEdit) BorderStyle() TBorderStyle {
    return LabeledEdit_GetBorderStyle(l.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (l *TLabeledEdit) SetBorderStyle(value TBorderStyle) {
    LabeledEdit_SetBorderStyle(l.instance, value)
}

func (l *TLabeledEdit) CharCase() TEditCharCase {
    return LabeledEdit_GetCharCase(l.instance)
}

func (l *TLabeledEdit) SetCharCase(value TEditCharCase) {
    LabeledEdit_SetCharCase(l.instance, value)
}

// 获取颜色。
//
// Get color.
func (l *TLabeledEdit) Color() TColor {
    return LabeledEdit_GetColor(l.instance)
}

// 设置颜色。
//
// Set color.
func (l *TLabeledEdit) SetColor(value TColor) {
    LabeledEdit_SetColor(l.instance, value)
}

// 获取约束控件大小。
func (l *TLabeledEdit) Constraints() *TSizeConstraints {
    return AsSizeConstraints(LabeledEdit_GetConstraints(l.instance))
}

// 设置约束控件大小。
func (l *TLabeledEdit) SetConstraints(value *TSizeConstraints) {
    LabeledEdit_SetConstraints(l.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (l *TLabeledEdit) DoubleBuffered() bool {
    return LabeledEdit_GetDoubleBuffered(l.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (l *TLabeledEdit) SetDoubleBuffered(value bool) {
    LabeledEdit_SetDoubleBuffered(l.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (l *TLabeledEdit) DragCursor() TCursor {
    return LabeledEdit_GetDragCursor(l.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (l *TLabeledEdit) SetDragCursor(value TCursor) {
    LabeledEdit_SetDragCursor(l.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (l *TLabeledEdit) DragMode() TDragMode {
    return LabeledEdit_GetDragMode(l.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (l *TLabeledEdit) SetDragMode(value TDragMode) {
    LabeledEdit_SetDragMode(l.instance, value)
}

func (l *TLabeledEdit) EditLabel() *TBoundLabel {
    return AsBoundLabel(LabeledEdit_GetEditLabel(l.instance))
}

// 获取控件启用。
//
// Get the control enabled.
func (l *TLabeledEdit) Enabled() bool {
    return LabeledEdit_GetEnabled(l.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (l *TLabeledEdit) SetEnabled(value bool) {
    LabeledEdit_SetEnabled(l.instance, value)
}

// 获取字体。
//
// Get Font.
func (l *TLabeledEdit) Font() *TFont {
    return AsFont(LabeledEdit_GetFont(l.instance))
}

// 设置字体。
//
// Set Font.
func (l *TLabeledEdit) SetFont(value *TFont) {
    LabeledEdit_SetFont(l.instance, CheckPtr(value))
}

// 获取隐藏选择。
func (l *TLabeledEdit) HideSelection() bool {
    return LabeledEdit_GetHideSelection(l.instance)
}

// 设置隐藏选择。
func (l *TLabeledEdit) SetHideSelection(value bool) {
    LabeledEdit_SetHideSelection(l.instance, value)
}

func (l *TLabeledEdit) LabelPosition() TLabelPosition {
    return LabeledEdit_GetLabelPosition(l.instance)
}

func (l *TLabeledEdit) SetLabelPosition(value TLabelPosition) {
    LabeledEdit_SetLabelPosition(l.instance, value)
}

func (l *TLabeledEdit) LabelSpacing() int32 {
    return LabeledEdit_GetLabelSpacing(l.instance)
}

func (l *TLabeledEdit) SetLabelSpacing(value int32) {
    LabeledEdit_SetLabelSpacing(l.instance, value)
}

// 获取最大长度。
func (l *TLabeledEdit) MaxLength() int32 {
    return LabeledEdit_GetMaxLength(l.instance)
}

// 设置最大长度。
func (l *TLabeledEdit) SetMaxLength(value int32) {
    LabeledEdit_SetMaxLength(l.instance, value)
}

// 获取只能输入数字。
func (l *TLabeledEdit) NumbersOnly() bool {
    return LabeledEdit_GetNumbersOnly(l.instance)
}

// 设置只能输入数字。
func (l *TLabeledEdit) SetNumbersOnly(value bool) {
    LabeledEdit_SetNumbersOnly(l.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (l *TLabeledEdit) ParentColor() bool {
    return LabeledEdit_GetParentColor(l.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (l *TLabeledEdit) SetParentColor(value bool) {
    LabeledEdit_SetParentColor(l.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (l *TLabeledEdit) ParentDoubleBuffered() bool {
    return LabeledEdit_GetParentDoubleBuffered(l.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (l *TLabeledEdit) SetParentDoubleBuffered(value bool) {
    LabeledEdit_SetParentDoubleBuffered(l.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (l *TLabeledEdit) ParentFont() bool {
    return LabeledEdit_GetParentFont(l.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (l *TLabeledEdit) SetParentFont(value bool) {
    LabeledEdit_SetParentFont(l.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (l *TLabeledEdit) ParentShowHint() bool {
    return LabeledEdit_GetParentShowHint(l.instance)
}

// 设置以父容器的ShowHint属性为准。
func (l *TLabeledEdit) SetParentShowHint(value bool) {
    LabeledEdit_SetParentShowHint(l.instance, value)
}

// 获取密码掩码字符。
func (l *TLabeledEdit) PasswordChar() uint16 {
    return LabeledEdit_GetPasswordChar(l.instance)
}

// 设置密码掩码字符。
func (l *TLabeledEdit) SetPasswordChar(value uint16) {
    LabeledEdit_SetPasswordChar(l.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (l *TLabeledEdit) PopupMenu() *TPopupMenu {
    return AsPopupMenu(LabeledEdit_GetPopupMenu(l.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (l *TLabeledEdit) SetPopupMenu(value IComponent) {
    LabeledEdit_SetPopupMenu(l.instance, CheckPtr(value))
}

// 获取只读。
func (l *TLabeledEdit) ReadOnly() bool {
    return LabeledEdit_GetReadOnly(l.instance)
}

// 设置只读。
func (l *TLabeledEdit) SetReadOnly(value bool) {
    LabeledEdit_SetReadOnly(l.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (l *TLabeledEdit) ShowHint() bool {
    return LabeledEdit_GetShowHint(l.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (l *TLabeledEdit) SetShowHint(value bool) {
    LabeledEdit_SetShowHint(l.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (l *TLabeledEdit) TabOrder() TTabOrder {
    return LabeledEdit_GetTabOrder(l.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (l *TLabeledEdit) SetTabOrder(value TTabOrder) {
    LabeledEdit_SetTabOrder(l.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (l *TLabeledEdit) TabStop() bool {
    return LabeledEdit_GetTabStop(l.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (l *TLabeledEdit) SetTabStop(value bool) {
    LabeledEdit_SetTabStop(l.instance, value)
}

// 获取文本。
func (l *TLabeledEdit) Text() string {
    strLen := l.GetTextLen()
    if strLen != 0 {
        var buffStr string
        l.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// 设置文本。
func (l *TLabeledEdit) SetText(value string) {
    LabeledEdit_SetText(l.instance, value)
}

// 获取提示文本。
func (l *TLabeledEdit) TextHint() string {
    return LabeledEdit_GetTextHint(l.instance)
}

// 设置提示文本。
func (l *TLabeledEdit) SetTextHint(value string) {
    LabeledEdit_SetTextHint(l.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (l *TLabeledEdit) Visible() bool {
    return LabeledEdit_GetVisible(l.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (l *TLabeledEdit) SetVisible(value bool) {
    LabeledEdit_SetVisible(l.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (l *TLabeledEdit) SetOnChange(fn TNotifyEvent) {
    LabeledEdit_SetOnChange(l.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (l *TLabeledEdit) SetOnClick(fn TNotifyEvent) {
    LabeledEdit_SetOnClick(l.instance, fn)
}

// 设置双击事件。
func (l *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
    LabeledEdit_SetOnDblClick(l.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (l *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
    LabeledEdit_SetOnDragDrop(l.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (l *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
    LabeledEdit_SetOnDragOver(l.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (l *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDrag(l.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (l *TLabeledEdit) SetOnEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnEnter(l.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (l *TLabeledEdit) SetOnExit(fn TNotifyEvent) {
    LabeledEdit_SetOnExit(l.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (l *TLabeledEdit) SetOnKeyDown(fn TKeyEvent) {
    LabeledEdit_SetOnKeyDown(l.instance, fn)
}

// 设置键键下事件。
func (l *TLabeledEdit) SetOnKeyPress(fn TKeyPressEvent) {
    LabeledEdit_SetOnKeyPress(l.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (l *TLabeledEdit) SetOnKeyUp(fn TKeyEvent) {
    LabeledEdit_SetOnKeyUp(l.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (l *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
    LabeledEdit_SetOnMouseDown(l.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (l *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseEnter(l.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (l *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseLeave(l.instance, fn)
}

// 设置鼠标移动事件。
func (l *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    LabeledEdit_SetOnMouseMove(l.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (l *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
    LabeledEdit_SetOnMouseUp(l.instance, fn)
}

// 获取能否撤销。
func (l *TLabeledEdit) CanUndo() bool {
    return LabeledEdit_GetCanUndo(l.instance)
}

// 获取修改。
//
// Get modified.
func (l *TLabeledEdit) Modified() bool {
    return LabeledEdit_GetModified(l.instance)
}

// 设置修改。
//
// Set modified.
func (l *TLabeledEdit) SetModified(value bool) {
    LabeledEdit_SetModified(l.instance, value)
}

// 获取选择的长度。
func (l *TLabeledEdit) SelLength() int32 {
    return LabeledEdit_GetSelLength(l.instance)
}

// 设置选择的长度。
func (l *TLabeledEdit) SetSelLength(value int32) {
    LabeledEdit_SetSelLength(l.instance, value)
}

// 获取选择的启始位置。
func (l *TLabeledEdit) SelStart() int32 {
    return LabeledEdit_GetSelStart(l.instance)
}

// 设置选择的启始位置。
func (l *TLabeledEdit) SetSelStart(value int32) {
    LabeledEdit_SetSelStart(l.instance, value)
}

// 获取选择的文本。
func (l *TLabeledEdit) SelText() string {
    return LabeledEdit_GetSelText(l.instance)
}

// 设置选择的文本。
func (l *TLabeledEdit) SetSelText(value string) {
    LabeledEdit_SetSelText(l.instance, value)
}

// 获取依靠客户端总数。
func (l *TLabeledEdit) DockClientCount() int32 {
    return LabeledEdit_GetDockClientCount(l.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (l *TLabeledEdit) DockSite() bool {
    return LabeledEdit_GetDockSite(l.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (l *TLabeledEdit) SetDockSite(value bool) {
    LabeledEdit_SetDockSite(l.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (l *TLabeledEdit) MouseInClient() bool {
    return LabeledEdit_GetMouseInClient(l.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (l *TLabeledEdit) VisibleDockClientCount() int32 {
    return LabeledEdit_GetVisibleDockClientCount(l.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (l *TLabeledEdit) Brush() *TBrush {
    return AsBrush(LabeledEdit_GetBrush(l.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (l *TLabeledEdit) ControlCount() int32 {
    return LabeledEdit_GetControlCount(l.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (l *TLabeledEdit) Handle() HWND {
    return LabeledEdit_GetHandle(l.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (l *TLabeledEdit) ParentWindow() HWND {
    return LabeledEdit_GetParentWindow(l.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (l *TLabeledEdit) SetParentWindow(value HWND) {
    LabeledEdit_SetParentWindow(l.instance, value)
}

func (l *TLabeledEdit) Showing() bool {
    return LabeledEdit_GetShowing(l.instance)
}

// 获取使用停靠管理。
func (l *TLabeledEdit) UseDockManager() bool {
    return LabeledEdit_GetUseDockManager(l.instance)
}

// 设置使用停靠管理。
func (l *TLabeledEdit) SetUseDockManager(value bool) {
    LabeledEdit_SetUseDockManager(l.instance, value)
}

func (l *TLabeledEdit) Action() *TAction {
    return AsAction(LabeledEdit_GetAction(l.instance))
}

func (l *TLabeledEdit) SetAction(value IComponent) {
    LabeledEdit_SetAction(l.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (l *TLabeledEdit) Align() TAlign {
    return LabeledEdit_GetAlign(l.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (l *TLabeledEdit) SetAlign(value TAlign) {
    LabeledEdit_SetAlign(l.instance, value)
}

func (l *TLabeledEdit) BoundsRect() TRect {
    return LabeledEdit_GetBoundsRect(l.instance)
}

func (l *TLabeledEdit) SetBoundsRect(value TRect) {
    LabeledEdit_SetBoundsRect(l.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (l *TLabeledEdit) ClientHeight() int32 {
    return LabeledEdit_GetClientHeight(l.instance)
}

// 设置客户区高度。
//
// Set client height.
func (l *TLabeledEdit) SetClientHeight(value int32) {
    LabeledEdit_SetClientHeight(l.instance, value)
}

func (l *TLabeledEdit) ClientOrigin() TPoint {
    return LabeledEdit_GetClientOrigin(l.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (l *TLabeledEdit) ClientRect() TRect {
    return LabeledEdit_GetClientRect(l.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (l *TLabeledEdit) ClientWidth() int32 {
    return LabeledEdit_GetClientWidth(l.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (l *TLabeledEdit) SetClientWidth(value int32) {
    LabeledEdit_SetClientWidth(l.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (l *TLabeledEdit) ControlState() TControlState {
    return LabeledEdit_GetControlState(l.instance)
}

// 设置控件状态。
//
// Set control state.
func (l *TLabeledEdit) SetControlState(value TControlState) {
    LabeledEdit_SetControlState(l.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (l *TLabeledEdit) ControlStyle() TControlStyle {
    return LabeledEdit_GetControlStyle(l.instance)
}

// 设置控件样式。
//
// Set control style.
func (l *TLabeledEdit) SetControlStyle(value TControlStyle) {
    LabeledEdit_SetControlStyle(l.instance, value)
}

func (l *TLabeledEdit) Floating() bool {
    return LabeledEdit_GetFloating(l.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (l *TLabeledEdit) Parent() *TWinControl {
    return AsWinControl(LabeledEdit_GetParent(l.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (l *TLabeledEdit) SetParent(value IWinControl) {
    LabeledEdit_SetParent(l.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (l *TLabeledEdit) Left() int32 {
    return LabeledEdit_GetLeft(l.instance)
}

// 设置左边位置。
//
// Set Left position.
func (l *TLabeledEdit) SetLeft(value int32) {
    LabeledEdit_SetLeft(l.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (l *TLabeledEdit) Top() int32 {
    return LabeledEdit_GetTop(l.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (l *TLabeledEdit) SetTop(value int32) {
    LabeledEdit_SetTop(l.instance, value)
}

// 获取宽度。
//
// Get width.
func (l *TLabeledEdit) Width() int32 {
    return LabeledEdit_GetWidth(l.instance)
}

// 设置宽度。
//
// Set width.
func (l *TLabeledEdit) SetWidth(value int32) {
    LabeledEdit_SetWidth(l.instance, value)
}

// 获取高度。
//
// Get height.
func (l *TLabeledEdit) Height() int32 {
    return LabeledEdit_GetHeight(l.instance)
}

// 设置高度。
//
// Set height.
func (l *TLabeledEdit) SetHeight(value int32) {
    LabeledEdit_SetHeight(l.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (l *TLabeledEdit) Cursor() TCursor {
    return LabeledEdit_GetCursor(l.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (l *TLabeledEdit) SetCursor(value TCursor) {
    LabeledEdit_SetCursor(l.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (l *TLabeledEdit) Hint() string {
    return LabeledEdit_GetHint(l.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (l *TLabeledEdit) SetHint(value string) {
    LabeledEdit_SetHint(l.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (l *TLabeledEdit) ComponentCount() int32 {
    return LabeledEdit_GetComponentCount(l.instance)
}

// 获取组件索引。
//
// Get component index.
func (l *TLabeledEdit) ComponentIndex() int32 {
    return LabeledEdit_GetComponentIndex(l.instance)
}

// 设置组件索引。
//
// Set component index.
func (l *TLabeledEdit) SetComponentIndex(value int32) {
    LabeledEdit_SetComponentIndex(l.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (l *TLabeledEdit) Owner() *TComponent {
    return AsComponent(LabeledEdit_GetOwner(l.instance))
}

// 获取组件名称。
//
// Get the component name.
func (l *TLabeledEdit) Name() string {
    return LabeledEdit_GetName(l.instance)
}

// 设置组件名称。
//
// Set the component name.
func (l *TLabeledEdit) SetName(value string) {
    LabeledEdit_SetName(l.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (l *TLabeledEdit) Tag() int {
    return LabeledEdit_GetTag(l.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (l *TLabeledEdit) SetTag(value int) {
    LabeledEdit_SetTag(l.instance, value)
}

// 获取左边锚点。
func (l *TLabeledEdit) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(LabeledEdit_GetAnchorSideLeft(l.instance))
}

// 设置左边锚点。
func (l *TLabeledEdit) SetAnchorSideLeft(value *TAnchorSide) {
    LabeledEdit_SetAnchorSideLeft(l.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (l *TLabeledEdit) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(LabeledEdit_GetAnchorSideTop(l.instance))
}

// 设置顶边锚点。
func (l *TLabeledEdit) SetAnchorSideTop(value *TAnchorSide) {
    LabeledEdit_SetAnchorSideTop(l.instance, CheckPtr(value))
}

// 获取右边锚点。
func (l *TLabeledEdit) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(LabeledEdit_GetAnchorSideRight(l.instance))
}

// 设置右边锚点。
func (l *TLabeledEdit) SetAnchorSideRight(value *TAnchorSide) {
    LabeledEdit_SetAnchorSideRight(l.instance, CheckPtr(value))
}

// 获取底边锚点。
func (l *TLabeledEdit) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(LabeledEdit_GetAnchorSideBottom(l.instance))
}

// 设置底边锚点。
func (l *TLabeledEdit) SetAnchorSideBottom(value *TAnchorSide) {
    LabeledEdit_SetAnchorSideBottom(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(LabeledEdit_GetChildSizing(l.instance))
}

func (l *TLabeledEdit) SetChildSizing(value *TControlChildSizing) {
    LabeledEdit_SetChildSizing(l.instance, CheckPtr(value))
}

// 获取边框间距。
func (l *TLabeledEdit) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(LabeledEdit_GetBorderSpacing(l.instance))
}

// 设置边框间距。
func (l *TLabeledEdit) SetBorderSpacing(value *TControlBorderSpacing) {
    LabeledEdit_SetBorderSpacing(l.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (l *TLabeledEdit) DockClients(Index int32) *TControl {
    return AsControl(LabeledEdit_GetDockClients(l.instance, Index))
}

// 获取指定索引子控件。
func (l *TLabeledEdit) Controls(Index int32) *TControl {
    return AsControl(LabeledEdit_GetControls(l.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (l *TLabeledEdit) Components(AIndex int32) *TComponent {
    return AsComponent(LabeledEdit_GetComponents(l.instance, AIndex))
}

// 获取锚侧面。
func (l *TLabeledEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(LabeledEdit_GetAnchorSide(l.instance, AKind))
}

